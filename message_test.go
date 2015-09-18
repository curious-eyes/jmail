package jmail

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"

	// "strings"
	// "fmt"
	// "golang.org/x/text/encoding/japanese"
	// "golang.org/x/text/transform"
	// "io/ioutil"
	// "mime"
)

func TestDecSubject(t *testing.T) {
	testemls := "./testsubj/"
	var f *os.File

	chksubj := []string{
		"Gophers at Gophercon",
		"【テスト環境】サイト更新が完了しました",
		"【テスト環境】サイト更新が完了しました",
		"【テスト環境】サイト更新が完了しました",
		"【テスト環境】サイト更新が完了しました",
	}
	// outstr, _ := utf8_to_2022(chksubj)
	// enc := mime.WordEncoder('b')
	// fmt.Println(enc.Encode("UTF-8", outstr))

	err := filepath.Walk(testemls,
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				// ディレクトリは無視
				rel, _ := filepath.Abs(path)
				// fmt.Println(rel)
				if f, err = os.Open(rel); err != nil {
					t.Errorf("test: Failed open file: %s (%v)", rel, err)
				}
				defer f.Close()

				var msg *Jmessage
				if msg, err = ReadMessage(f); err != nil {
					t.Errorf("test: ReadMessage error: %s (%v)", rel, err)
				}
				filen := info.Name()
				seq, _ := strconv.Atoi(filen[0:2])
				if msg.DecSubject() != chksubj[seq] {
					t.Errorf("test: Subject error: %s (%s)", filen, msg.DecSubject())
				}
			}
			return nil
		})
	if err != nil {
		t.Errorf("test: Failed filepath.Walk: %v", err)
	}

}

func TestDecBody(t *testing.T) {
	testemls := "./testbody/"
	var f *os.File

	chkbody := []string{
		"Message body\r\n",
		"サイトを更新した状態に保つことはセキュリティにとって重要です。それはまた、あなたとあなたの読者にとってインターネットをより安全な場所にすることでもあります。\r\n",
		"サイトを更新した状態に保つことはセキュリティにとって重要です。それはまた、あなたとあなたの読者にとってインターネットをより安全な場所にすることでもあります。\r\n",
		"サイトを更新した状態に保つことはセキュリティにとって重要です。それはまた、あなたとあなたの読者にとってインターネットをより安全な場所にすることでもあります。\n",
		"サイトを更新した状態に保つことはセキュリティにとって重要です。それはまた、あなたとあなたの読者にとってインターネットをより安全な場所にすることでもあります。\r\n",
		"go go gopher!\r\n",
		"サイトを更新した状態に保つことはセキュリティにとって重要です。それはまた、あなたとあなたの読者にとってインターネットをより安全な場所にすることでもあります。[image:\r\ntalks.png][image: doc.png]\r\n",
	}

	err := filepath.Walk(testemls,
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				// ディレクトリは無視
				rel, _ := filepath.Abs(path)
				// fmt.Println(rel)
				if f, err = os.Open(rel); err != nil {
					t.Errorf("test: Failed open file: %s (%v)", rel, err)
				}
				defer f.Close()

				var msg *Jmessage
				if msg, err = ReadMessage(f); err != nil {
					t.Errorf("test: ReadMessage error: %s (%v)", rel, err)
				}
				filen := info.Name()
				seq, _ := strconv.Atoi(filen[0:2])
				var body []byte
				if body, err = msg.DecBody(); err != nil {
					t.Errorf("test: Body error: %s (%v)", filen, err)
					return err
				}
				if string(body) != chkbody[seq] {
					t.Errorf("test: Body error: %s (%s)", filen, chkbody[seq])
				}
			}
			return nil
		})
	if err != nil {
		t.Errorf("test: Failed filepath.Walk: %v", err)
	}

}

// // UTF-8 から ISO-2022-JP
// func utf8_to_2022(str string) (string, error) {
//   iostr := strings.NewReader(str)
//   rio := transform.NewReader(iostr, japanese.ISO2022JP.NewEncoder())
//   ret, err := ioutil.ReadAll(rio)
//   if err != nil {
//           return "", err
//   }
//   return string(ret), err
// }
