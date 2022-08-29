package boluobao

import (
	"fmt"
	url_ "net/url"
	"sf/cfg"
	req "sf/src/https"
	_struct "sf/struct"
	"sf/struct/sfacg_structs"
	"sf/struct/sfacg_structs/bookshelf"
	"strconv"
)

func GET_BOOK_INFORMATION(NovelId string) (_struct.Books, error) {
	params := map[string]string{"expand": "intro,tags,sysTags,totalNeedFireMoney,originTotalNeedFireMoney"}
	response := req.Get("novels/"+NovelId, &sfacg_structs.BookInfo{}, params).(*sfacg_structs.BookInfo)
	if response.Status.HTTPCode == 200 && response.Data.NovelName != "" {
		return _struct.Books{
			NovelName:  cfg.RegexpName(response.Data.NovelName),
			NovelID:    strconv.Itoa(response.Data.NovelID),
			NovelCover: response.Data.NovelCover,
			AuthorName: response.Data.AuthorName,
			CharCount:  strconv.Itoa(response.Data.CharCount),
			MarkCount:  strconv.Itoa(response.Data.MarkCount),
			SignStatus: response.Data.SignStatus,
		}, nil
	} else {
		return _struct.Books{}, fmt.Errorf(response.Status.Msg.(string))
	}

}

func GET_ACCOUNT_INFORMATION() *sfacg_structs.Account {
	return req.Get("user", &sfacg_structs.Account{}, nil).(*sfacg_structs.Account)
}

func GET_BOOK_SHELF_INFORMATION() *bookshelf.InfoData {
	params := map[string]string{"expand": "novels,albums,comics,discount"}
	return req.Get("user/Pockets", &bookshelf.InfoData{}, params).(*bookshelf.InfoData)
}

func GET_CATALOGUE(NovelID string) []map[string]string {
	var division_info []map[string]string
	var chapter_index int
	response := req.Get(fmt.Sprintf("novels/%v/dirs", NovelID), &sfacg_structs.Catalogue{}, map[string]string{"expand": "originNeedFireMoney"})
	for division_index, division := range response.(*sfacg_structs.Catalogue).Data.VolumeList {
		fmt.Printf("第%v卷\t\t%v\n", division_index+1, division.Title)
		for _, chapter := range division.ChapterList {
			chapter_index += 1
			division_info = append(division_info, map[string]string{
				"division_name":  division.Title,
				"division_id":    strconv.Itoa(division.VolumeID),
				"division_index": strconv.Itoa(division_index),
				"chapter_name":   chapter.Title,
				"chapter_id":     strconv.Itoa(chapter.ChapID),
				"chapter_index":  strconv.Itoa(chapter_index),
				"money":          strconv.Itoa(chapter.OriginNeedFireMoney),
				"file_name":      cfg.Config_file_name(division_index, chapter_index, strconv.Itoa(chapter.ChapID)),
			})
		}
	}
	return division_info

}

func GET_CONTENT(cid string) *sfacg_structs.Content {
	params := map[string]string{"expand": "content"}
	if result := req.Get("Chaps/"+cid, &sfacg_structs.Content{}, params); result != nil {
		return result.(*sfacg_structs.Content)
	} else {
		return GET_CONTENT(cid) // retry once if failed to get content
	}
}

func GET_SEARCH(keyword string, page int) *sfacg_structs.Search {
	params := map[string]string{"q": url_.QueryEscape(keyword), "size": "20", "page": strconv.Itoa(page)}
	return req.Get("search/novels/result", &sfacg_structs.Search{}, params).(*sfacg_structs.Search)

}

func LOGIN_ACCOUNT(username, password string) *sfacg_structs.Login {
	params := fmt.Sprintf(`{"username":"%s", "password": "%s"}`, username, password)
	response, Cookie := req.Login(req.SET_URL("sessions", nil), []byte(params))
	for _, cookie := range Cookie {
		response.Cookie += cookie.Name + "=" + cookie.Value + ";"
	}
	return response
}
