#include "cholesky.hxx"
#include "daal.h"
#include <iostream>

using namespace daal;
using namespace daal::algorithms;
using namespace daal::data_management;
using namespace daal::services;

int choleskyDecompose(int dimension, double inputArray[]) {

    /* Create input numeric table from array */
    SharedPtr<NumericTable> inputData = SharedPtr<NumericTable>(new Matrix<double>(dimension, dimension, inputArray));

    /*  Create the algorithm object for computation of the Cholesky decomposition using the default method */
    cholesky::Batch<> algorithm;

    /* Set input for the algorithm */
    algorithm.input.set(cholesky::data, inputData);

    /* Compute Cholesky decomposition */
    algorithm.compute();

    /* Get pointer to Cholesky factor */
    SharedPtr<Matrix<double> > factor =
        staticPointerCast<Matrix<double>, NumericTable>(algorithm.getResult()->get(cholesky::choleskyFactor));

    /* Return the first element of the Cholesky factor */
    return (*factor)[0][0];
}
