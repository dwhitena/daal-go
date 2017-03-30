#ifndef CHOLESKY_H
#define CHOLESKY_H

// __cplusplus gets defined when a C++ compiler processes the file.
// extern "C" is needed so the C++ compiler exports the symbols w/out name issues.
#ifdef __cplusplus
extern "C" {
#endif

int choleskyDecompose(int dimension, double inputArray[]);

#ifdef __cplusplus
}
#endif


#endif
