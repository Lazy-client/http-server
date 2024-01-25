package main

//
//import (
//"fmt"
//"sync"
//"testing"
//"time"
//
//"tencent.com/apps/fc/ifbook/srcs/test/runner/logic"
//"tencent.com/apps/fc/ifbook/srcs/test/runner/logic/executor/extends"
//"tencent.com/protocols/fc/ifbook/interface/adapter"
//interfacePb "tencent.com/protocols/fc/ifbook/interface/core"
//pb "tencent.com/protocols/fc/ifbook/test/core"
//
//"google.golang.org/protobuf/types/known/structpb"
//)
//
//var workspace = "test"
//
//func run() {
//	workspace, _ := logic.CreateWorkspace()
//	runner := logic.HTTPInterfaceCaseRunner{
//		RootWorkspace: workspace,
//		Workspace:     workspace,
//		InterfaceCaseRequest: &pb.RunInterfaceCaseRequest{
//			Case: &pb.InterfaceCase{
//				Spec: &pb.InterfaceCaseSpec{
//					InterfaceRef: &interfacePb.Interface{
//						Category: interfacePb.InterfaceCategory_INTERFACE_CATEGORY_HTTP,
//					},
//					// 断言
//					Assertions: []*pb.Assertion{
//						{
//							Disabled: false,
//							Expression: &pb.AssertExpression{
//								Expressions: []*pb.AssertExpression{
//									{
//										ActualValue: &pb.AssertExpression_ActualValue{
//											Ref:   pb.RefLocation_REF_SCRIPT,
//											Value: "print(123)",
//										},
//									},
//								},
//							},
//						},
//						{
//							Disabled: false,
//							Expression: &pb.AssertExpression{
//								Expressions: []*pb.AssertExpression{
//									{
//										ActualValue: &pb.AssertExpression_ActualValue{
//											Ref:   pb.RefLocation_REF_HEADER,
//											Value: "Content-Type",
//										},
//										ExpectValue: "application/json",
//										Operator:    pb.AssertExpression_INCLUDE,
//										ValueType:   pb.FieldType_FIELD_TYPE_STRING,
//									},
//								},
//							},
//						},
//					},
//					// 请求
//					Request: &pb.InterfaceCaseSpec_Request{
//						Settings: &pb.InterfaceCaseSpec_Request_Settings{
//							EncodeUrlAutomatically: true,
//						},
//						Url: &pb.InterfaceCaseSpec_Request_HttpURL{
//							Uri: "",
//						},
//						RequestBody: &adapter.HttpInterface_Request_RequestBody{
//							Type:      adapter.HttpInterface_Request_RequestBody_RAW,
//							RawFormat: adapter.HttpInterface_Request_RequestBody_TEXT,
//							Content:   structpb.NewStringValue(`{"Case":{"Meta":{"DescName":"response是二进制流","Priority":"P2","InterfaceUid":"ifbif-rP2wZgNKQC","RepoName":"ifbook-managed","RepoVersionName":"baseline","Namespace":"5646"},"Spec":{"Request":{"Url":{"Scheme":"HTTP","Uri":"/","Variables":[],"QueryParams":[]},"Method":"GET","Headers":[{"Name":"PRIVATE-TOKEN","Value":"OdTPUMEkmg06Zxc3uP9w","Type":"string","Required":false,"Desc":"","ContentType":""}],"Cookies":[],"Authorization":{"Type":"NO_AUTH","Props":[],"Key":"","Value":"","TargetLocation":""},"RequestBody":null,"TargetInstance":{"Name":"http://sanweishuwu.tech","Type":"dns"},"Settings":{"InsecureSkipVerify":false,"EncodeURLAutomatically":true,"AutomaticallyFollowRedirects":true,"FollowOriginalHttpMethod":true,"FollowAuthorizationHeader":true,"FollowRedirectsMaximumTimes":10,"RemoveRefererHeaderOnRedirect":false,"KeepUnknownVariables":false,"IgnoreUnknownVariables":false}},"PreOperations":[{"Name":"未命名","Type":"CUSTOM_SCRIPT","Disabled":false,"Params":{"Language":"LANGUAGE_PYTHON3","Script":"print(123)\n\n"}}],"PostOperations":[{"Name":"未命名","Type":"CUSTOM_SCRIPT","Disabled":true,"Params":{"Language":"LANGUAGE_PYTHON3","Script":"// 判断响应码是否为固定值\npm.test(\"Status code is 200\", function () {\n    pm.response.to.have.status(200);\n    console.log(123);\n});\n\n"}}],"Assertions":[{"Name":"test","Disabled":false,"Expression":{"ExpectValue":"","ActualValue":{"Ref":"REF_SCRIPT","Value":"with ifbook.test(\"Json key equals value\") as tc:\n    jsonData = ifbook.response.json();\n    tc.assertEqual(jsonData.value, 100)\n\nwith ifbook.test(\"Body matches string\") as tc:\n    tc.assertIn(\"string_you_want_to_search\", ifbook.response.text())\n"}}}]}},"Env":{"Meta":{"Uid":""}},"Runner":{"Labels":["IDC - 内置节点"]}}`),
//						},
//						Target: "http://www.baidu.com",
//						Headers: []*pb.FieldProperty{
//							{
//								Name:  "test2",
//								Value: "{{test1}}",
//							},
//						},
//					},
//				},
//			},
//			TearDowns: []*pb.TearDown{},
//		},
//
//		RuntimeContext: &extends.RuntimeContext{
//			Globals: []extends.Variable{
//				{
//					Key: "test1",
//					Value: &structpb.Value{
//						Kind: &structpb.Value_StringValue{
//							StringValue: "value1",
//						},
//					},
//				},
//			},
//		},
//	}
//	runner.Run()
//	wg.Done()
//}
//
//var wg sync.WaitGroup
//
//func BenchmarkRunner(t *testing.B) {
//
//	//body := bytes.NewBufferString(`{"code":200}`)
//	//mock := gomonkey.ApplyMethodReturn(&http.Client{}, "Do",
//	//	&http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(body), Header: http.Header{"Content-Type": {"application/json"}}}, nil)
//	//defer mock.Reset()
//	now := time.Now()
//	for i := 0; i < t.N; i++ {
//		wg.Add(1)
//		go run()
//	}
//	wg.Wait()
//	fmt.Println(time.Since(now).Seconds(), "s")
//}
//
