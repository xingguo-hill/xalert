package alert

/**
https://open.feishu.cn/document/client-docs/bot-v3/add-custom-bot
*/
import (
	"fmt"
	"os/exec"
	"strconv"
)

func SendFeishuContentAlert(url, content string) {
	params := []string{
		"-X", "POST",
		"-H", "Content-Type: application/json",
		"-d", fmt.Sprintf(`{"msg_type":"text","content":{"text":%s}}`, strconv.Quote(content)),
	}
	cmd := exec.Command("curl", append([]string{url}, params...)...)
	cmd.Stdout = nil
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing curl:", err)
		return
	}
}
