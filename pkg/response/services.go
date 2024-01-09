package response

import "net/http"

var BadRequestResponse = StructResponse{
	HttpStatus: http.StatusBadRequest,
	StateCode:  "BAD_REQUEST",
	MessageEn:  "Invalid request payload",
	MessageBn:  "অবৈধ অনুরোধ পেইলোড",
}

var ValidationErrorResponse = StructResponse{
	HttpStatus: http.StatusBadRequest,
	StateCode:  "VALIDATION_ERROR",
	MessageEn:  "Validation error",
	MessageBn:  "মান যাচাই করার ত্রুটি",
}

var InternalServerErrorResponse = StructResponse{
	HttpStatus: http.StatusInternalServerError,
	StateCode:  "INTERNAL_SERVER_ERROR",
	MessageEn:  "Failed",
	MessageBn:  "ব্যর্থ",
}

var NotFoundResponse = StructResponse{
	HttpStatus: http.StatusNotFound,
	StateCode:  "NOT_FOUND",
	MessageEn:  "Service not found",
	MessageBn:  "সার্ভিস পাওয়া যায়নি",
}

var ServiceCreateResponse = StructResponse{
	HttpStatus: http.StatusCreated,
	StateCode:  "SUCCESS",
	MessageEn:  "Service created successfully",
	MessageBn:  "সার্ভিস সফলভাবে তৈরি হয়েছে",
}

var GetServicesResponse = StructResponse{
	HttpStatus: http.StatusOK,
	StateCode:  "SUCCESS",
	MessageEn:  "Services fetched successfully",
	MessageBn:  "সার্ভিসগুলো সফলভাবে পাওয়া গিয়েছে",
}

var GetServiceByNameResponse = StructResponse{
	HttpStatus: http.StatusOK,
	StateCode:  "SUCCESS",
	MessageEn:  "Service fetched successfully",
	MessageBn:  "সার্ভিস সফলভাবে পাওয়া গিয়েছে",
}

var UpdateServiceResponse = StructResponse{
	HttpStatus: http.StatusOK,
	StateCode:  "SUCCESS",
	MessageEn:  "Service updated successfully",
	MessageBn:  "সার্ভিস সফলভাবে আপডেট হয়েছে",
}

var DeleteServiceResponse = StructResponse{
	HttpStatus: http.StatusOK,
	StateCode:  "SUCCESS",
	MessageEn:  "Service deleted successfully",
	MessageBn:  "সার্ভিস সফলভাবে মুছে ফেলা হয়েছে",
}
