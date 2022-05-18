package mailservice

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/i-rm/wb/be/pb"
)

type Post interface {
	Confirm(url, email, pass string) *pb.MailConfirmResponse
	Reset(url, email string) *pb.ResetResponse
}

// DBStore ...
type PostFix struct {
	mutex sync.Mutex
}

// NewDBStore ...
func NewPostFix() *PostFix {
	return &PostFix{
		mutex: sync.Mutex{},
	}
}

func call(command string) {
	cmd := exec.Command("sudo", "bash", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	err1 := cmd.Wait()
	if err1 != nil {
		fmt.Println(err1)
	}
}

func (post *PostFix) Confirm(url, email, pass string) *pb.MailConfirmResponse {
	post.mutex.Lock()
	defer post.mutex.Unlock()

	body := `"<p>Поздравляем с регистрацией на <b><a href=''>www.bdrop.com</a></b></p><p>Чтобы подтвердить регистрацию, пожалуйста, откройте ссылку:</p><a href='` + url + `'>Подтвердить регистрацию</a><p>После подтверждения можно использовать следующие данные для входа на сайт:</p><table style='text-align:left'><tbody><tr><th>Email:</th><td>` + email + `</td></tr><tr><th>Пароль:</th><td>` + pass + `</td></tr></tbody></table>"`
	subj := `"$(echo -e "Подтверждение регистрации\nFrom: www.bdrop.net <noreply@bdrop.net>\nContent-Type: text/html")"`
	call(`echo ` + body + ` | mail -s ` + subj + ` ` + email + ``)
	resp := &pb.MailConfirmResponse{
		StatusCode: http.StatusOK,
	}
	return resp
}

func (post *PostFix) Reset(url, email string) *pb.ResetResponse {
	post.mutex.Lock()
	defer post.mutex.Unlock()

	body := `"<p>Мы получили запрос на восстановление доступа к Вашему аккаунту.</p><p>Если этот запрос сделан Вами, пожалуйста, откройте ссылку для установки нового пароля:</p><a href='` + url + `'>` + url + `</a>"`
	subj := `"$(echo -e "Восстановление доступа к Вашему аккаунту\nFrom: www.bdrop.net <noreply@bdrop.net>\nContent-Type: text/html")"`
	call(`echo ` + body + ` | mail -s ` + subj + ` ` + email + ``)
	resp := &pb.ResetResponse{
		StatusCode: http.StatusOK,
	}
	return resp
}
