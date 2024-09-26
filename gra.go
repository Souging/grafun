package main

import (
	"bytes"
	//"crypto/ecdsa"
	//"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"
	"strings"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	//"github.com/ethereum/go-ethereum/crypto"
)

var (
	addr  string
	miyao string
	hasd  string
	cs    int
)

func openBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Printf("Error opening browser: %v\n", err)
		fmt.Printf("Please open the following URL in your browser: %s\n", url)
	}

	return err
}

func creat() {

	response, err := http.Get("http://199.119.138.181:9988/addr")
	if err != nil {
        fmt.Println("Error:", err)
        return
    }
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Error reading body:", err)
        return
    }
	result := string(body)
	splitResult := strings.Split(result, "|")
	if len(splitResult) >= 2 {
		addr = splitResult[0]
		miyao = splitResult[1]
	}else{
		creat()
	}
	//fmt.Printf("%s|%s\n", addr, miyao)
}

func hashde() {
	postData := fmt.Sprintf(`{"sign":"","referrer":"0x3bE4D49447b927a3aFC82C8ED43Ac58993a91712","wallet_name":"HOT","wallet_address":"%s"}`, addr)
	
	//fmt.Println(postData)
	req, err := http.NewRequest("POST", "https://wl-api.gra.fun/api/hash_addr", bytes.NewBuffer([]byte(postData)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	hashedWalletAddress, ok := result["hashed_wallet_address"].(string)
	if !ok {
		fmt.Println("Failed to get hashed_wallet_address")
		return
	}

	hasd = hashedWalletAddress
}

func bsc(){

	postData := fmt.Sprintf(`{"jsonrpc":"2.0","id":0,"method":"eth_getBalance","params":["%s","latest"]}`, addr)
	req, err := http.NewRequest("POST", "https://rpc.ankr.com/bsc", bytes.NewBuffer([]byte(postData)))
	if err != nil {
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	resp.Body.Close()

}

func status() {
	times := 0
	for {
		req, err := http.NewRequest("GET", "https://wl-api.gra.fun/api/status/"+addr, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}
		if times >= 100{
			break
		}
		var result map[string]interface{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			//fmt.Println("Error parsing JSON:", err)
			times++
			time.Sleep(1 * time.Second)
			continue
		}
		break	
	}

	// Send YouTube task
	req, err := http.NewRequest("POST", "https://wl-api.gra.fun/api/check_youtube/"+addr, nil)
	if err != nil {
		fmt.Println("Error creating YouTube task request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending YouTube task request:", err)
		return
	}
	resp.Body.Close()
	postData := fmt.Sprintf(`{"jsonrpc":"2.0","id":1,"method":"eth_getBalance","params":["%s","latest"]}`, addr)
	req, err = http.NewRequest("POST", "https://rpc.ankr.com/bsc", bytes.NewBuffer([]byte(postData)))
	if err != nil {
		return
	}
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	resp.Body.Close()
	// Query coin status
	req, err = http.NewRequest("GET", "https://wl-api.gra.fun/api/status/"+addr, nil)
	if err != nil {
		//fmt.Println("Error creating coin status request:", err)
		return
	}

	resp, err = client.Do(req)
	if err != nil {
		//fmt.Println("Error sending coin status request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println("Error reading coin status response:", err)
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		//fmt.Println("Error parsing coin status JSON:", err)
		return
	}
	coin := result["offchain_points"]

	postData = fmt.Sprintf(`{"jsonrpc":"2.0","id":2,"method":"eth_getBalance","params":["%s","latest"]}`, addr)
	req, err = http.NewRequest("POST", "https://rpc.ankr.com/bsc", bytes.NewBuffer([]byte(postData)))
	if err != nil {
		return
	}
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	resp.Body.Close()
	
	
	
	cs++
	fmt.Printf("%s|%s|%v|%d\n", addr, miyao, coin, cs)
	
}
func main() {
	for i := 1; i <= 100000; i++ {
		creat()
		bsc()
		hashde()
		//str := "tg://resolve?domain=GraFunBot&start=" + hasd
		fmt.Println("/start ",hasd)
		//return
		//err := openBrowser(str)
		//if err != nil {
			//fmt.Println("Error opening browser:", err)
		//}
		status()
	}
}
