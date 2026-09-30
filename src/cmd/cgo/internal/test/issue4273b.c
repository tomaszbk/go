#ifdef __ELF__
extern void _compilerrt_abort_impl(const char *file, int line, const char *func);

void __my_abort(const char *file, int line, const char *func) {
	_compilerrt_abort_impl(file, line, func);
}
#endif
