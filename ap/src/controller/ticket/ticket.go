package ticket

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/util/dateutil"
	"alma-server/ap/src/common/util/htmlutil"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/event/EventRpcService"
	"alma-server/ap/src/domain/menu/MenuService"
	"alma-server/ap/src/domain/ticket/TicketRpcService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"alma-server/ap/src/infrastructure/grpc/proto/ticket"
	"net/http"
)

// PageHTML チケットの画面
func PageHTML(w http.ResponseWriter, r *http.Request) {

	// param
	req := &common.EventRequest{
		Event: param.Value(r, "event"),
	}

	ctx := r.Context()
	mid := almactx.GetMid(ctx)
	txTime := almactx.GetTxTime(ctx)

	result := TicketRpcService.Page(ctx, mid, txTime, req.Event)

	response.BaseHTML(
		w,
		"チケット",
		"/template/controller/ticket/head.html",
		map[string]interface{}{},
		"/template/controller/ticket/ticket.html",
		map[string]interface{}{
			"ticketInfoList": result.TicketInfoList,
			"eventId":        result.EventId,
		},
		[]string{
			"/static/js/common/table/table.js",
			"/static/js/controller/ticket/ticket.js",
		},
		[]string{
			"/static/css/common/table/table.css",
			"/static/css/common/content_head_button/content_head_button.css",
			"/static/css/controller/ticket/ticket.css",
		},
		result.EventName,
		MenuService.GetMenu("ticket_top", "ticket"),
	)
}

// CreatePageHTML チケットの作成画面
func CreatePageHTML(w http.ResponseWriter, r *http.Request) {

	// param
	req := &common.EventRequest{
		Event: param.Value(r, "event"),
	}

	ctx := r.Context()
	mid := almactx.GetMid(ctx)

	result := EventRpcService.GetEvent(ctx, mid, req.Event)

	response.BaseHTML(
		w,
		"チケット作成",
		"/template/controller/ticket/create/head.html",
		map[string]interface{}{},
		"/template/controller/ticket/create/ticket_create.html",
		map[string]interface{}{
			"ticketForm": htmlutil.CreateTemplateToString(
				"/template/component/ticket/form.html",
				map[string]interface{}{},
			),
		},
		[]string{
			"/static/js/component/ticket/form.js",
			"/static/js/controller/ticket/create/ticket_create.js",
		},
		[]string{
			"/static/css/common/content_head_button/content_head_button.css",
			"/static/css/component/ticket/form.css",
			"/static/css/controller/ticket/create/ticket_create.css",
		},
		result.EventName,
		MenuService.GetMenu("ticket_top", ""),
	)
}

// UpdatePageHTML チケットの編集画面
func UpdatePageHTML(w http.ResponseWriter, r *http.Request) {

	// param
	req := &ticket.UpdatePageRequest{
		EventId:  param.Value(r, "event"),
		TicketId: param.Value(r, "ticketId"),
	}

	ctx := r.Context()
	mid := almactx.GetMid(ctx)
	txTime := almactx.GetTxTime(ctx)

	result := TicketRpcService.UpdatePage(ctx, mid, txTime, req.EventId, req.TicketId)

	response.BaseHTML(
		w,
		"チケットの編集",
		"/template/controller/ticket/update/head.html",
		map[string]interface{}{},
		"/template/controller/ticket/update/ticket_update.html",
		map[string]interface{}{
			"ticketForm": htmlutil.CreateTemplateToString(
				"/template/component/ticket/form.html",
				map[string]interface{}{
					"ticketId":    result.TicketInfo.TicketId,
					"ticketName":  result.TicketInfo.TicketName,
					"ticketPrice": result.TicketInfo.TicketPrice,
					"ticketDesc":  result.TicketInfo.TicketDesc,
				},
			),
		},
		[]string{
			"/static/js/component/ticket/form.js",
			"/static/js/controller/ticket/update/ticket_update.js",
		},
		[]string{
			"/static/css/common/content_head_button/content_head_button.css",
			"/static/css/component/ticket/form.css",
			"/static/css/controller/ticket/update/ticket_update.css",
		},
		result.EventName,
		MenuService.GetMenu("ticket_top", ""),
	)

}

// CreateTicket チケットの作成
func CreateTicket(w http.ResponseWriter, r *http.Request) {

	// param
	req := &ticket.CreateTicketRequest{}
	err := param.JSON(r, req)
	chk.SE(err)

	ctx := r.Context()
	mid := almactx.GetMid(ctx)
	txTime := almactx.GetTxTime(ctx)

	TicketRpcService.CreateTicket(ctx, mid, txTime,
		req.TicketInfo.EventId, req.TicketInfo.TicketId,
		req.TicketInfo.TicketName, req.TicketInfo.TicketPrice,
		req.TicketInfo.TicketDesc, req.TicketInfo.TicketStock,
		dateutil.TimestampToTime(req.TicketInfo.TicketEventStartTime),
	)

	response.JSON(w, &common.Empty{})
}

// UpdateTicket チケット情報の更新
func UpdateTicket(w http.ResponseWriter, r *http.Request) {

	req := &ticket.UpdateTicketRequest{}
	err := param.JSON(r, req)
	chk.SE(err)

	ctx := r.Context()
	mid := almactx.GetMid(ctx)
	txTime := almactx.GetTxTime(ctx)

	TicketRpcService.UpdateTicket(ctx, mid, txTime, req.EventId, req.BeforeTicketId, req.TicketInfo)

	response.JSON(w, &common.Empty{})
}

// DeleteTicket チケットの削除
// TODO ログを残すか、UseYnで削除するようにする
func DeleteTicket(w http.ResponseWriter, r *http.Request) {

	req := &ticket.DeleteTicketRequest{}
	err := param.JSON(r, req)
	chk.SE(err)

	ctx := r.Context()
	mid := almactx.GetMid(ctx)
	txTime := almactx.GetTxTime(ctx)

	// Delete
	TicketRpcService.DeleteTicket(ctx, mid, txTime, req.EventId, req.TicketIdList)

	response.JSON(w, &common.Empty{})
}
