package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//u := "http://www.sanweishuwu.tech/event-hook/send"
	//u := "http://www.sanweishuwu.tech/event-hook/a"
	//u := "http://www.sanweishuwu.tech/event-hook/b"
	u := "http://127.0.0.1/event-hook/send"
	method := http.MethodGet
	/*payload := strings.NewReader(`{
	    "event": {
	        "metadata": {
	            "creationTimestamp": "2023-08-31T07:09:22.425989327Z",
	            "namespace": "7130",
	            "uid": "21d0e610-4037-407b-bafe-e27c0a2d4ecb"
	        },
	        "payload": {
	            "@type": "type.googleapis.com/core.InterfaceEventPayload",
	            "Resource": {
	                "ActionType": "OPERATION_ACTION_UPDATE",
	                "RepoName": "ifbook-managed",
	                "RepoVersionName": "baseline",
	                "ResourceType": "NODE_TYPE_INTERFACE",
	                "Uid": "ifbif-jBiI44ZnuQ",
	                "UidPath": "ifbpf-pmgiEMKy7X,ifbif-jBiI44ZnuQ"
	            }
	        }
	    },
	    "eventType": {
	        "allowSubscriptionTypes": [
	            "SUB_TYPE_INTERNAL",
	            "SUB_TYPE_EXTERNAL"
	        ],
	        "metadata": {
	            "name": "ifbook.interface.change"
	        }
	    },
	    "metadata": {
	        "creationTimestamp": "2023-08-31T07:09:22.430433830Z",
	        "uid": "997e0d4d-9995-423d-9272-78940c0edcd6"
	    }
	}`)*/

	req, err := http.NewRequest(method, u, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("token", "445c584ea89358488414aa6f8a81dede")
	req.Header.Add("Host", "account.billing.tencentyun.com")

	client := &http.Client{}
	req.Host = "account.billing.tencentyun.com"
	resp, err := client.Do(req)
	content, _ := io.ReadAll(resp.Body)

	log.Println(content)
	log.Println(req.URL.String())
	log.Println(resp.Request.URL.String())
}
