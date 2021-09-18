package urls

import bgsave "github.com/YarikRevich/HideSeek-Services/pkg/bg_save"

func PinURLToServices(host string){
	bgsave.UseService().SetURL(host)
};