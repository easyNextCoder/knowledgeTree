%module rules_cgo
%include "std_vector.i"

%{
#include "cgo_wrap.h"
%}

%include "cgo_wrap.h"
namespace std{
    %template(IntVector) vector<int>;
}
