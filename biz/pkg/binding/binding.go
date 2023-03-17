package binding

import "github.com/cloudwego/hertz/pkg/app/server/binding"

// https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/binding-and-validate/#%E9%85%8D%E7%BD%AE%E5%85%B6%E4%BB%96-json-unmarshal-%E5%BA%93

func init() {
	// 使用标准库
	binding.UseStdJSONUnmarshaler()

	//// 使用gjson
	//binding.UseGJSONUnmarshaler()
	//
	//// 使用第三方json unmarshal方法
	//binding.UseThirdPartyJSONUnmarshaler()
}
