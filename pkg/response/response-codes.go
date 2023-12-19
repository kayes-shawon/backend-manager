package response

import "net/http"

var TechnicalError = StructResponse{
	HttpStatus: http.StatusBadRequest,
	StateCode:  "BM_TE_400",
	MessageEn:  "Due to a technical error, we are unable to process further. Please try again later.",
	MessageBn:  "টেকনিক্যাল ত্রুটি ঘটায় পরবর্তী ধাপে প্রসেস করা সম্ভব হচ্ছে না। অনুগ্রহ করে পরে আবার চেষ্টা করুন।",
}
