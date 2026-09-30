//go:build ignore

#ifdef WIN32
#if defined(EXPORT_DLL)
#    define VAR __declspec(dllexport)
#elif defined(IMPORT_DLL)
#    define VAR __declspec(dllimport)
#endif
#else
#    define VAR extern
#endif

VAR const char *exported_var;
