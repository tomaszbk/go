package issue48082

import "init" /* ERROR "init must be a func" */ /* ERROR "could not import init" */
