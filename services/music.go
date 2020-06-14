package services

func PlayYoutube(link string, guild string, channel string) {
	// video, err := ytdl.GetVideoInfo(link)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return // Returning to avoid crash when video informations could not be found
	// }

	// for _, format := range video.Formats {
	// 	if format.AudioEncoding == "opus" || format.AudioEncoding == "aac" || format.AudioEncoding == "vorbis" {
	// 		data, err := video.GetDownloadURL(format)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		url := data.String()
	// 		// go playAudioFile(url, guild, channel, "youtube")
	// 		return
	// 	}
	// }

}
