package main

/*

Written by Jack Watts for SMPTE Standards TC 27C Community.

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

import (
	"flag"
	"fmt"
	"os"

	"github.com/jack-watts/empty-tt/pkg/tt"
)

func main() {
	run()
}

// run available command line flags.
func run() {

	flag.BoolVar(&tt.Txt, "text", true, "- Inidcate that text profile is to be used.")
	flag.BoolVar(&tt.Img, "image", false, "- Inidcate that image profile is to be used.")
	flag.BoolVar(&tt.Track, "T", false, "- write MXF trackfile, requires '-d'")
	flag.BoolVar(&tt.Encrypt, "e", false, "- encrypt trackfile")
	flag.IntVar(&tt.Duration, "d", 24, "- set the duration of the track file.")
	flag.StringVar(&tt.Framerate, "p", "24", "- set the frame rate of the track file.")
	flag.IntVar(&tt.Display, "m", 0, "- set the DisplayType.'0'=MainSubtitle,'1'=ClosedCaption. (default '0')")
	flag.IntVar(&tt.Reel, "r", 1, "- set the ReelNumber, Default ='1'")
	flag.StringVar(&tt.Language, "l", "en", "- set the RFC 5646 Language subtag")
	flag.StringVar(&tt.Title, "t", "No Title", "- set the ContentTitleText value.")
	flag.StringVar(&tt.Template, "x", "", "- path to 428-7 XML to use as template")
	flag.StringVar(&tt.Output, "o", "", "- set the output path, Default is StdOut")
	flag.Parse()
	if len(flag.Args()) > 0 {
		fmt.Println("check command expression")
		os.Exit(1)
	}
	if err := tt.CreateXML(tt.Txt, tt.Img, tt.Track, tt.Encrypt, tt.Reel, tt.Display, tt.Duration, tt.Framerate, tt.Language, tt.Title, tt.Template, tt.Output); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return
}
