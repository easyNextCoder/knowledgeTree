����
ͷ�ļ�	����
<algorithm>	�㷨����
<numeric>	��ֵ�㷨
<functional>	��������/�º���
����
No.	����	˵��
1	�ǿɱ������㷨	Non-modifying sequence operations	��ֱ���޸��������ݵ��㷨��
2	�ɱ������㷨	Modifying sequence operations	�����޸��������ݵ��㷨��
3	�����㷨	Sorting/Partitions/Binary search/	���������򡢺ϲ��������㷨������
4	��ֵ�㷨	Merge/Heap/Min/max	���������ݽ�����ֵ���㡣
���
����	����
fill(beg,end,val)	��ֵval����[beg,end)��Χ�ڵ�����Ԫ�ء�
fill_n(beg,n,val)	��ֵval����[beg,beg+n)��Χ�ڵ�����Ԫ�ء�
generate(beg,end,func)	�������ú���func���[beg,end)��Χ�ڵ�����Ԫ�ء�
generate_n(beg,n,func)	�������ú���func���[beg,beg+n)��Χ�ڵ�����Ԫ�ء�
fill()/fill_n()���������ֵͬ��generate()/generate_n()������䲻ֵͬ��
����/�任
����	����
for_each(beg,end,func)	��[beg,end)��Χ������Ԫ�����ε��ú���func������func�����޸������е�Ԫ�ء�
transform(beg,end,res,unary)	��[beg,end)��Χ������Ԫ�����ε��ú���unary���������res�С�
transform(beg2,end1,beg2,res,binary)	��[beg,end)��Χ������Ԫ����[beg2,beg2+end-beg)������Ԫ�����ε��ú���unary���������res�С�
�����С
����	����
max(a,b)	��������Ԫ���нϴ�һ����
max(a,b,cmp)	ʹ���Զ���Ƚϲ���cmp,��������Ԫ���нϴ�һ����
max_element(beg,end)	����һ��ForwardIterator��ָ��[beg,end)������Ԫ�ء�
max_element(beg,end,cmp)	ʹ���Զ���Ƚϲ���cmp,����һ��ForwardIterator��ָ��[beg,end)������Ԫ�ء�
min(a,b)	��������Ԫ���н�Сһ����
min(a,b,cmp)	ʹ���Զ���Ƚϲ���cmp,��������Ԫ���н�Сһ����
min_element(beg,end)	����һ��ForwardIterator��ָ��[beg,end)����С��Ԫ�ء�
min_element(beg,end,cmp)	ʹ���Զ���Ƚϲ���cmp,����һ��ForwardIterator��ָ��[beg,end)����С��Ԫ�ء�
�����㷨(12��)���ṩԪ���������
1 ����
����	����
sort(beg,end)	Ĭ��������������Ԫ��
sort(beg,end,comp)	ʹ�ú���comp����Ƚϲ�����ִ��sort()��
partition(beg,end,pred)	Ԫ����������ʹ��pred�������ѽ��Ϊtrue��Ԫ�ط��ڽ��Ϊfalse��Ԫ��֮ǰ��
stable_sort(beg,end)	��sort()���ƣ��������Ԫ��֮���˳���ϵ��
stable_sort(beg,end,pred)	ʹ�ú���pred����Ƚϲ�����ִ��stable_sort()��
stable_partition(beg,end)	��partition()���ƣ����������е����˳��
stable_partition(beg,end,pred)	ʹ�ú���pred����Ƚϲ�����ִ��stable_partition()��
partial_sort(beg,mid,end)	�������򣬱�����Ԫ�ظ����ŵ�[beg,end)�ڡ�
partial_sort(beg,mid,end,comp)	ʹ�ú���comp����Ƚϲ�����ִ��partial_sort()��
partial_sort_copy(beg1,end1,beg2,end2)	��partial_sort()���ƣ�ֻ�ǽ�[beg1,end1)��������и��Ƶ�[beg2,end2)��
partial_sort_copy(beg1,end1,beg2,end2,comp)	ʹ�ú���comp����Ƚϲ�����ִ��partial_sort_copy()��
nth_element(beg,nth,end)	����Ԫ��������������ʹ����С�ڵ�n��Ԫ�ص�Ԫ�ض���������ǰ�棬���������Ķ������ں��档
nth_element(beg,nth,end,comp)	ʹ�ú���comp����Ƚϲ�����ִ��nth_element()��
2 ��ת/��ת
����	����
reverse(beg,end)	Ԫ�����·�������
reverse_copy(beg,end,res)	��reverse()���ƣ����д��res��
rotate(beg,mid,end)	Ԫ���Ƶ�����ĩβ����mid��Ϊ������һ��Ԫ�ء�
rotate_copy(beg,mid,end,res)	��rotate()���ƣ����д��res
3 ���
����	����
random_shuffle(beg,end)	Ԫ�������������
random_shuffle(beg,end,gen)	ʹ�ú���gen����������ɺ���ִ��random_shuffle()��
�����㷨(13��)���ж��������Ƿ����ĳ��ֵ
1 ͳ��
����	����
count(beg,end,val)	����==����������[beg,end)��Ԫ����val���бȽϣ��������Ԫ�ظ�����
count_if(beg,end,pred)	ʹ�ú���pred����==������ִ��count()��
2 ����
����	����
find(beg,end,val)	����==����������[beg,end)��Ԫ����val���бȽϡ���ƥ��ʱ�������������ظ�Ԫ�ص�InputIterator��
find_if(beg,end,pred)	ʹ�ú���pred����==������ִ��find()��
find_first_of(beg1,end1,beg2,end2)	��[beg1,end1)��Χ�ڲ���[beg2,end2)������һ��Ԫ�صĵ�һ�γ��֡����ظ�Ԫ�ص�Iterator��
find_first_of(beg1,end1,beg2,end2,pred)	ʹ�ú���pred����==������ִ��find_first_of()�����ظ�Ԫ�ص�Iterator��
find_end(beg1,end1,beg2,end2)	��[beg1,end1)��Χ�ڲ���[beg2,end2)���һ�γ��֡��ҵ��򷵻����һ�Եĵ�һ��ForwardIterator�����򷵻�end1��
find_end(beg1,end1,beg2,end2,pred)	ʹ�ú���pred����==������ִ��find_end()�����ظ�Ԫ�ص�Iterator��
adjacent_find(beg,end)	��[beg,end)��Ԫ�أ�����һ�������ظ�Ԫ�أ��ҵ��򷵻�ָ�����Ԫ�صĵ�һ��Ԫ�ص�ForwardIterator�����򷵻�end��
adjacent_find(beg,end,pred)	ʹ�ú���pred����==������ִ��adjacent_find()��
3 ����
����	����
search(beg1,end1,beg2,end2)	��[beg1,end1)��Χ�ڲ���[beg2,end2)��һ�γ��֣�����һ��ForwardIterator�����ҳɹ�,����[beg1,end1)�ڵ�һ�γ���[beg2,end2)��λ�ã�����ʧ��ָ��end1��
search(beg1,end1,beg2,end2,pred)	ʹ�ú���pred����==������ִ��search()��
search_n(beg,end,n,val)	��[beg,end)��Χ�ڲ���val����n�ε�������
search_n(beg,end,n,val,pred)	ʹ�ú���pred����==������ִ��search_n()��
binary_search(beg,end,val)	��[beg,end)�в���val���ҵ�����true��
binary_search(beg,end,val,comp)	ʹ�ú���comp����Ƚϲ�����ִ��binary_search()��
4 �߽�
����	����
lower_bound(beg,end,val)	��[beg,end)��Χ�ڵĿ��Բ���val�����ƻ�����˳��ĵ�һ��λ�ã�����һ��ForwardIterator��
lower_bound(beg,end,val,comp)	ʹ�ú���comp����Ƚϲ�����ִ��lower_bound()��
upper_bound(beg,end,val)	��[beg,end)��Χ�ڲ���val�����ƻ�����˳������һ��λ�ã���λ�ñ�־һ������val��ֵ������һ��ForwardIterator��
upper_bound(beg,end,val,comp)	ʹ�ú���comp����Ƚϲ�����ִ��upper_bound()��
equal_range(beg,end,val)	����һ��iterator����һ����ʾlower_bound���ڶ�����ʾupper_bound��
equal_range(beg,end,val,comp)	ʹ�ú���comp����Ƚϲ�����ִ��lower_bound()��
ɾ�����滻�㷨(15��)
����
����	����
copy(beg,end,res)	����[beg,end)��res
copy_backward(beg,end,res)	��copy()��ͬ������Ԫ�������෴˳�򱻿�����
�Ƴ�
����	����
remove(beg,end,val)	ɾ��[beg,end)�����е���val��Ԫ�ء�ע�⣬�ú�����������ɾ��������
remove_if(beg,end,pred)	ɾ��[beg,end)��pred���Ϊtrue��Ԫ�ء�
remove_copy(beg,end,res,val)	�����в�����valԪ�ظ��Ƶ�res������OutputIteratorָ�򱻿�����ĩԪ�ص���һ��λ�á�
remove_copy_if(beg,end,res,pred)	������ʹpred���Ϊtrue��Ԫ�ؿ�����res��
�滻
����	����
replace(beg,end,oval,nval)	��[beg,end)�����е���oval��Ԫ�ض���nval���档
replace_copy(beg,end,res,oval,nval)	��replace()���ƣ����������д��res��
replace_if(beg,end,pred,nval)	��[beg,end)������predΪtrue��Ԫ����nval���档
replace_copy_if(beg,end,res,pred,nval)	��replace_if()�����������д��res��
ȥ��
����	����
unique(beg,end)	��������������ظ�Ԫ�أ���������ɾ��Ԫ�ء����ذ汾ʹ���Զ���Ƚϲ�����
unique(beg,end,pred)	������ʹpred���Ϊtrue�������ظ�Ԫ��ȥ�ء�
unique_copy(beg,end,res)	��unique���ƣ������ѽ�������res��
unique_copy(beg,end,res,pred)	��unique���ƣ������ѽ�������res��
����
����	����
swap(a,b)	�����洢��a��b�е�ֵ��
swap_range(beg1,end1,beg2)	��[beg1,end1)�ڵ�Ԫ��[beg2,beg2+beg1-end1)Ԫ��ֵ���н�����
iter_swap(it_a,it_b)	��������ForwardIterator��ֵ��
�����㷨(4��)<numeric>
����	����
accumulate(beg,end,val)	��[beg,end)��Ԫ��֮�ͣ��ӵ���ʼֵval�ϡ�
accumulate(beg,end,val,binary)	������binary����ӷ����㣬ִ��accumulate()��
partial_sum(beg,end,res)	��[beg,end)�ڸ�λ��ǰ����Ԫ��֮�ͷŽ�res�С�
partial_sum(beg,end,res,binary)	������binary����ӷ����㣬ִ��partial_sum()��
adjacent_difference(beg1,end1,res)	��[beg,end)��ÿ����ֵ����ǰԪ������һ��Ԫ�صĲ�Ž�res�С�
adjacent_difference(beg1,end1,res,binary)	������binary����������㣬ִ��adjacent_difference()��
inner_product(beg1,end1,beg2,val)	�������������ڻ�(��ӦԪ����ˣ������)�����ڻ��ӵ���ʼֵval�ϡ�
inner_product(beg1,end1,beg2,val,binary1,binary2)	������binary1����ӷ�����,��binary2����˷����㣬ִ��inner_product()��
��ϵ�㷨(4��)
����	����
equal(beg1,end1,beg2)	�ж�[beg1,end1)��[beg2,end2)��Ԫ�ض����
equal(beg1,end1,beg2,pred)	ʹ��pred��������Ĭ�ϵ�==��������
includes(beg1,end1,beg2,end2)	�ж�[beg1,end1)�Ƿ����[beg2,end2)��ʹ�õײ�Ԫ�ص�<���������ɹ�����true�����ذ汾ʹ���û�����ĺ�����
includes(beg1,end1,beg2,end2,comp)	������comp����<��������ִ��includes()��
lexicographical_compare(beg1,end1,beg2,end2)	���ֵ����ж�[beg1,end1)�Ƿ�С��[beg2,end2)
lexicographical_compare(beg1,end1,beg2,end2,comp)	������comp����<��������ִ��lexicographical_compare()��
mismatch(beg1,end1,beg2)	���бȽ�[beg1,end1)��[beg2,end2)��ָ����һ����ƥ���λ�ã�����һ��iterator����־��һ����ƥ��Ԫ��λ�á������ƥ�䣬����ÿ��������end��
mismatch(beg1,end1,beg2,pred)	ʹ��pred��������Ĭ�ϵ�==��������
�����㷨(6��)
����	����
merge(beg1,end1,beg2,end2,res)	�ϲ�[beg1,end1)��[beg2,end2)��ŵ�res��
merge(beg1,end1,beg2,end2,res,comp)	������comp����<��������ִ��merge()��
inplace_merge(beg,mid,end)	�ϲ�[beg,mid)��[mid,end)���������[beg,end)��
inplace_merge(beg,mid,end,cmp)	������comp����<��������ִ��inplace_merge()��
set_union(beg1,end1,beg2,end2,res)	ȡ[beg1,end1)��[beg2,end2)Ԫ�ز�����ŵ�res��
set_union(beg1,end1,beg2,end2,res,comp)	������comp����<��������ִ��set_union()��
set_intersection(beg1,end1,beg2,end2,res)	ȡ[beg1,end1)��[beg2,end2)Ԫ�ؽ�����ŵ�res��
set_intersection(beg1,end1,beg2,end2,res,comp)	������comp����<��������ִ��set_intersection()��
set_difference(beg1,end1,beg2,end2,res)	ȡ[beg1,end1)��[beg2,end2)Ԫ���ڲ��ŵ�res��
set_difference(beg1,end1,beg2,end2,res,comp)	������comp����<��������ִ��set_difference()��
set_symmetric_difference(beg1,end1,beg2,end2,res)	ȡ[beg1,end1)��[beg2,end2)Ԫ������ŵ�res��
��������㷨(2��)���ṩ����������ϰ�һ��˳������п����������
����	����
next_permutation(beg,end)	ȡ��[beg,end)�ڵ�����һ�����С�
next_permutation(beg,end,comp)	������comp����<��������ִ��next_permutation()��
prev_permutation(beg,end)	ȡ��[beg,end)�ڵ�����һ�����С�
prev_permutation(beg,end,comp)	������comp����<��������ִ��prev_permutation()��
���㷨(4��)
����	����
make_heap(beg,end)	��[beg,end)�ڵ�Ԫ������һ���ѡ�
make_heap(beg,end,comp)	������comp����<��������ִ��make_heap()��
pop_heap(beg,end)	��������ѡ�����first��last-1������Ȼ����������һ���ѡ���ʹ��������back�����ʱ�"����"��Ԫ�ػ���ʹ��pop_back����������ɾ�����������������Ԫ�شӶ��е�����
pop_heap(beg,end,comp)	������comp����<��������ִ��pop_heap()��
push_heap(beg,end)	����first��last-1��һ����Ч�ѣ�Ҫ�����뵽�ѵ�Ԫ�ش����λ��last-1���������ɶѡ���ָ��ú���ǰ�������Ȱ�Ԫ�ز���������
push_heap(beg,end,comp)	������comp����<��������ִ��push_heap()��
sort_heap(beg,end)	��[beg,end)�ڵ�������������
sort_heap(beg,end,comp)	������comp����<��������ִ��push_heap()��

���ߣ�jdzhangxin
���ӣ�https://www.jianshu.com/p/eb554b0943ab
��Դ������
����Ȩ���������С���ҵת������ϵ���߻����Ȩ������ҵת����ע��������
