/*
 * Dummy implementations for Windows, because Windows doesn't
 * support Unix-style signal handling.
 */

int install_handler() {
	return 0;
}


int check_handler() {
	return 0;
}
