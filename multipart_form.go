package http

import "github.com/Laky-64/http/types"

type multiPartOption types.MultiPartInfo

func (ct multiPartOption) Apply(o *types.RequestOptions) {
	tmpMultiPartInfo := types.MultiPartInfo(ct)
	o.MultiPart = &tmpMultiPartInfo
}

func MultiPartForm(data map[string]string, files map[string]types.FileDescriptor) RequestOption {
	return multiPartOption(types.MultiPartInfo{
		Data:  data,
		Files: files,
	})
}
