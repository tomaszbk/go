// This .cc file will be automatically compiled by the go tool and
// included in the package.

#include <string>
#include "main.h"

std::string Caller::call() {
	if (callback_ != 0)
		return callback_->run();
	return "";
}
