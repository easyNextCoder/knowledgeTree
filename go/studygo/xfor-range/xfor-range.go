package xfor_range

//1. 关于for-range中的坑
//https://zhuanlan.zhihu.com/p/105435646
/*
	for-range编译之后的源代码
  // The loop we generate:
  //   len_temp := len(range)
  //   range_temp := range//会把原来的arr拷贝，所以便利arr很浪费内存，便利slice和map则比较高效
  //   for index_temp = 0; index_temp < len_temp; index_temp++ {
  //           value_temp = range_temp[index_temp]//这个的item也发生了值拷贝
  //           index = index_temp
  //           value = value_temp
  //           original body
  //   }
*/
