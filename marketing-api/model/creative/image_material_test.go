package creative

import (
	"encoding/json"
	"testing"
)

func TestImageMaterialMarshaler(t *testing.T) {
	encs := []string{`{"image_info":[{"image_id":"test"}]}`, `{"image_info":{"image_id":"test"}}`, `{}`}
	objs := []ImageMaterial{
		{
			ImageInfo: &ImageInfoWrapper{
				List: []ImageInfo{
					{
						ImageID: "test",
					},
				},
			},
		},
		{
			ImageInfo: &ImageInfoWrapper{
				Object: &ImageInfo{
					ImageID: "test",
				},
			},
		},
		{
			ImageInfo: nil,
		},
	}
	for idx, enc := range encs {
		if bs, err := json.Marshal(objs[idx]); err != nil {
			t.Error(err)
		} else if string(bs) != enc {
			t.Errorf("expect:%s, got:%s", enc, string(bs))
		}
		var obj ImageMaterial
		if err := json.Unmarshal([]byte(enc), &obj); err != nil {
			t.Error(err)
		} else if objs[idx].ImageInfo != nil && obj.ImageInfo == nil || objs[idx].ImageInfo == nil && obj.ImageInfo != nil {
			t.Errorf("case:%s, expect: %+v, got: %+v\n", enc, objs[idx], obj)
		} else if objs[idx].ImageInfo != nil && obj.ImageInfo != nil && objs[idx].ImageInfo.IsObject() != obj.ImageInfo.IsObject() {
			t.Errorf("case: %s, expect: %+v, got: %+v\n", enc, objs[idx].ImageInfo, obj.ImageInfo)
		}
	}
}
