// Package tt enables a user to create a minimal ST 428-7 Subtitle
// XML Document in accordance with ISDCF Technicial Doc 16 - SMPTE
// ST 428-7 D-Cinema Distribution Master Subtitle - Minimal Empty Document Requirements
// as per said requirements stipulated in RDD 52 - SMPTE DCP Bv2.1 Application Profile
// available at https://doi.org/10.5594/SMPTE.RDD52.2020.
//
/* Copyright (c) 2020, Jack Watts. All rights reserved.

This program is free software : you can redistribute it and / or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT,
INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE

You should have received a copy of the GNU General Public License
along with this program.If not, see <http://www.gnu.org/licenses/>.*/

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
	as          = "asdcp-wrap"
)

// The proceeding list of exported variables are to be used when calling either CreateXML()
// or CreateXMF() functions.
var (
	// Text identifies that the Subtitle Text profile is to be used.
	Txt bool
	// Image identifies that the Subtitle Image profile is to be used.
	Img bool
	// Track signals that an MXF track file is to be created when using the CreateXML() function.
	Track bool
	// Encrypt is to be used in accordance with 'Track' as it signals that the resulting MXF is
	// to be encrypted.
	Encrypt bool
	// Reel is the reel number and shall be a positive integer reflecting the reel number the XML is to be used for
	Reel int
	// Duration is a positive integer value that maps to the ContainerDuration entry of the resulting MXF track file.
	Duration int
	// Display identifies what DisplayType value to be used. 0 = MainSubtitle, >= 1 = ClosedCaption.
	Display int
	// Framerate results in the EditRate of the Subtitle XML file and also translates to the TimeCodeRate element.
	Framerate string
	// Template is to be used when wanting to use an existing XML document to template the XML's general properties.
	Template string
	// Title is the value that populates the ContentTitleText element.
	Title string
	// Language is the RFC 5646 compliant subtag as per the IANA subtag registry.
	Language string
	// Output is the target output directory.
	Output string
	// unexported variables
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
	fontPath   string = getFont()
	mxfFileExt string = "_sub.mxf"
)

// ================================
// Begin exported functions

// CreateXML creates a St 428-7 compliant minimal XML document.
func CreateXML(Txt, Img, Track, Encrypt bool, Reel, Display, Duration int, FrameRate, Language, Title, Template, Output string) error {
	var subElement *Subtitle

	if Template != "" {
		s, err := parseXML(Template)
		if err != nil {
			fmt.Printf("%s\nunable to use template, running with default values\n", err)
		} else {
			xmlNs = s.XMLName.Space
			Title = s.ContentTitleText
			Language = s.Language
			Framerate = s.EditRate[:2]
			timecodeRate = s.TimeCodeRate
			if s.DisplayType == "MainSubtitle" {
				Display = 0
			}
			if s.DisplayType == "ClosedCaption" {
				Display = 1
				mxfFileExt = "_cap.mxf"
			}
		}
	}
	dxml := SubtitleReel{
		Xmlns:            xmlNs,
		ID:               "urn:uuid:" + docID,
		ContentTitleText: Title,
		IssueDate:        issueDate.Format(time.RFC3339)[:19] + "-00:00",
		ReelNumber:       Reel,
		Language:         Language,
		EditRate:         Framerate + " 1",
		TimeCodeRate:     Framerate,
		StartTime:        startTime,
		SubtitleList:     &Font{},
	}

	if Display == 0 {
		dxml.DisplayType = "MainSubtitle"
	}
	if Display >= 1 {
		dxml.DisplayType = "ClosedCaption"
		mxfFileExt = "_cap.mxf"
	}

	if Img {
		Txt = false
		imageID := makePNG(Output)
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
	if Txt {
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
	filename := docID + reelNo + strconv.Itoa(Reel) + xmlFileExt

	if enc, err := xml.MarshalIndent(dxml, "", "  "); err == nil {
		enc = []byte(xml.Header + string(enc))

		// Write XML to StdOut
		if Output == "" {
			fmt.Printf("%s\n", enc)
			return nil
		}
		if Output != "" {
			// Write out files to given output path.
			xmlOutputPath := filepath.Join(Output, filename)
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
			if Txt {
				fontFileName, err := filepath.Abs(fontPath)
				if err != nil {
					fmt.Println(err)
					fmt.Println("unable to resolve default font resource")
				}
				if err := copy(fontFileName, Output); err != nil {
					fmt.Println(err)
				}
			}

			// Handle Track File writing
			if Track {
				if testBinary() {
					if err := CreateMXF(Encrypt, Framerate, Output, xmlOutputPath, Reel, Duration); err != nil {
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
// It expects asdcp-wrap to be present at yout system's $PATH.
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
		filename := filepath.Join(output, ID)
		f, err := os.Create(filename)
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
