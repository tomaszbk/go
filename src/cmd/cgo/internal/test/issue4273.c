#ifdef __ELF__
__attribute__((weak))
__attribute__((visibility("hidden")))
void _compilerrt_abort_impl(const char *file, int line, const char *func) {
}
#endif
