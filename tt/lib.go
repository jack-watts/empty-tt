// Package tt ...
package tt

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	fontName    = "232c45d8-fde8-4e5e-86b9-86e96354daf3"
	defaultFont = "resources/font/" + fontName
	startTime   = "00:00:00:00"
	timeIn      = "00:00:04:00"
	timeOut     = "00:00:19:00"
	reelNo      = "_r"
	xmlFileExt  = ".xml"
	mxfFileExt  = "_sub.mxf"
	as          = "asdcp-wrap"
)

var (
	docID         string    = uuidType4()
	mxfID         string    = uuidType4()
	issueDate     time.Time = time.Now()
	timecodeRate  string
	dcst2007      string = "http://www.smpte-ra.org/schemas/428-7/2007/DCST"
	dcst2010      string = "dcst2010"
	dcst2014      string = "dcst2014"
	xmlNs         string = "http://www.smpte-ra.org/schemas/428-7/2014/DCST"
	xmlNsSubtitle        = map[string]string{
		"http://www.smpte-ra.org/schemas/428-7/2010/DCST": dcst2010,
		"http://www.smpte-ra.org/schemas/428-7/2014/DCST": dcst2014,
	}
	fontPath string = getFont()
)

// ================================
// Begin exported functions

// CreateXML creates a St 428-7 compliant minimal XML document.
func CreateXML(text, image, track, encrypt bool, reel, display, duration int, frameRate, lang, title, template, output string) error {
	var subElement *Subtitle

	if template != "" {
		s, err := parseXML(template)
		if err != nil {
			fmt.Printf("%s\nunable to use template, running with default values\n", err)
		} else {
			xmlNs = s.XMLName.Space
			title = s.ContentTitleText
			lang = s.Language
			frameRate = s.EditRate[:2]
			timecodeRate = s.TimeCodeRate
			if s.DisplayType == "MainSubtitle" {
				display = 0
			}
			if s.DisplayType == "ClosedCaption" {
				display = 1
			}
		}
	}
	dxml := SubtitleReel{
		Xmlns:            xmlNs,
		ID:               "urn:uuid:" + docID,
		ContentTitleText: title,
		IssueDate:        issueDate.Format(time.RFC3339)[:19] + "-00:00",
		ReelNumber:       reel,
		Language:         lang,
		EditRate:         frameRate + " 1",
		TimeCodeRate:     frameRate,
		StartTime:        startTime,
		SubtitleList:     &Font{},
	}

	if display == 0 {
		dxml.DisplayType = "MainSubtitle"
	}
	if display >= 1 {
		dxml.DisplayType = "ClosedCaption"
	}

	if image {
		text = false
		imageID := makePNG(output)
		subElement = &Subtitle{
			TimeIn:  timeIn,
			TimeOut: timeOut,
			Image: []*Image{
				&Image{
					Image: "urn:uuid:" + imageID,
				},
			},
		}
	}
	if text {
		dxml.LoadFont = &LoadFont{
			ID:   "Arial",
			Font: fontName,
		}
		subElement = &Subtitle{
			TimeIn:  timeIn,
			TimeOut: timeOut,
			Text: []*Text{
				&Text{
					Text: "",
				},
			},
		}
	}
	dxml.SubtitleList.Subtitle = append(dxml.SubtitleList.Subtitle, subElement)

	// Begin XML Writing
	filename := docID + reelNo + strconv.Itoa(reel) + xmlFileExt

	if enc, err := xml.MarshalIndent(dxml, "", "  "); err == nil {
		enc = []byte(xml.Header + string(enc))

		// Write XML to StdOut
		if output == "" {
			fmt.Printf("%s\n", enc)
			return nil
		}
		if output != "" {
			// Write out files to given output path.
			xmlOutputPath := filepath.Join(output, filename)
			f, err := os.Create(xmlOutputPath)
			if err != nil {
				fmt.Println(err)
			}
			if _, err := f.Write(enc); err != nil {
				f.Close() // Close file and return error.
				fmt.Println(err)
			}
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}

			// Copy Font to output path
			fontFileName, err := filepath.Abs(fontPath)
			if err != nil {
				fmt.Println(err)
				fmt.Println("unable to resolve default font resource")
			}
			if err := copy(fontFileName, output); err != nil {
				fmt.Println(err)
			}

			// Handle Track File writing
			if track {
				if testBinary() {
					if err := CreateMXF(encrypt, frameRate, output, xmlOutputPath, reel, duration); err != nil {
						fmt.Println("error writing MXF")
						return err
					}
					return nil
				}
				fmt.Println("asdcp not installed or not available at $PATH")
			}
		}

	}
	return nil
}

// CreateMXF creates a D-Cinema track file using pre-defined asdcp-wrap command line arguments.
// It expects asdcp-wrap to be present at $PATH
func CreateMXF(encrypt bool, frameRate, output, filename string, reel, duration int) error {
	mxfFilename := mxfID + reelNo + strconv.Itoa(reel) + mxfFileExt
	outputPath := filepath.Join(output, mxfFilename)
	asArgs := []string{}
	keyID := uuidType4()
	keyStr := randomHex()

	if encrypt {
		asArgs = []string{"-L", "-j", keyID, "-k", keyStr, "-a", mxfID, "-d", strconv.Itoa(duration), "-p", frameRate, filename, outputPath}
		fmt.Printf(`
Keep the following safe!
KeyID: %s
KeyString: %s
`, keyID, keyStr)
	}
	if !encrypt {
		asArgs = []string{"-L", "-a", mxfID, "-d", strconv.Itoa(duration), "-p", frameRate, filename, outputPath}
	}
	cmd := exec.Command(as, asArgs...)
	if _, err := cmd.Output(); err != nil {
		return err
	}
	return nil
}

// End exported functions

// ================================
// Begin unexported functions

// parseXML parsing a given ST 428-7 XML document to use the the document's global properties in the newly created document.
func parseXML(filename string) (*SubtitleReel, error) {
	var s *SubtitleReel
	file, err := filepath.Abs(filename)
	if err != nil {
		fmt.Println("given file path is not absolute:", err)
	}
	if path.Ext(file) == ".xml" {
		f, err := os.Open(file)
		defer f.Close()
		if err != nil {
			fmt.Println(err)
			return s, err
		}
		bytestream, _ := ioutil.ReadAll(f)
		xml.Unmarshal(bytestream, &s)
		for xmlns := range xmlNsSubtitle {
			if xmlns == s.XMLName.Space {
				return s, nil
			}
			if xmlns == dcst2007 {
				err := errors.New("invalid namespace")
				fmt.Printf("%s: %s\n", err, dcst2007)
				return s, err
			}
		}
	}
	err = errors.New("template document type cannot be determined, returning nil value")
	return s, err
}

// makePNG ...
func makePNG(output string) string {
	const width, height = 128, 128
	ID := uuidType4()

	if output != "" {
		// Create a transparent image of the given width and height.
		img := image.NewRGBA(image.Rect(0, 0, width, height))

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				img.Set(x, y, image.Transparent)
			}
		}

		f, err := os.Create(ID)
		if err != nil {
			log.Fatal(err)
		}

		if err := png.Encode(f, img); err != nil {
			f.Close()
			log.Fatal(err)
		}

		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
	return ID
}

// uuidType4 generates a canonical string representation of a Type-4 UUID.
func uuidType4() string {
	u := uuid.NewV4()
	return u.String()
}

// randomHex returns a random 16 byte hex string.
func randomHex() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// getFont returns the default Font resource
func getFont() string {
	exe, _ := os.Executable()
	dir, _ := filepath.Split(filepath.Dir(exe))
	return filepath.Join(dir, defaultFont)
}

// copy facilitates a standard OS copy function
func copy(file, dst string) error {
	// stat the file
	stat, err := os.Stat(file)
	if err != nil {
		fmt.Println("unable to stat file: ", file)
		return err
	}

	// confirm file is a regular file
	if !stat.Mode().IsRegular() {
		fmt.Println("not regular file:", file)
		return nil
	}

	// read the file
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("unable to read file: ", stat.Name())
		return err
	}
	defer f.Close()
	// create the destination
	filepath := filepath.Join(dst, stat.Name())
	target, err := os.Create(filepath)
	if err != nil {
		fmt.Println("unable to read destination: ", dst)
		return err
	}
	defer target.Close()

	// write the file to destination
	if _, err := io.Copy(target, f); err != nil {
		fmt.Println("copy process interupted: ", err)
		return err
	}
	return nil
}

// testBinary performs a boolean assesment of the asdcp-wrap binary at $PATH
func testBinary() bool {
	args := []string{as}

	if runtime.GOOS == "windows" {
		_, err := exec.Command("where", args...).Output()
		if err != nil {
			return false
		}
		return true
	}

	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		_, err := exec.Command("which", args...).Output()
		if err != nil {
			return false
		}
		return true
	}

	return false
}

// End unexported functions
