// +build !static

package git

/*
#include <git2.h>
#cgo LDFLAGS: -I/usr/local/Cellar/libgit2/0.28.3/include -L/usr/local/Cellar/libgit2/0.28.3/lib -lgit2

#if LIBGIT2_VER_MAJOR != 0 || LIBGIT2_VER_MINOR != 28
# error "Invalid libgit2 version; this git2go supports libgit2 v0.28"
#endif

*/
import "C"
