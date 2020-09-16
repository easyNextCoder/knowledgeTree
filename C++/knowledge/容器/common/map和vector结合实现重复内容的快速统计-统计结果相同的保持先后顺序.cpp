对于题目中的要求一种实现方法是：
vector getAll()
map insert(vector)
vector2 get(map)
sort(vector)

上面的方法可以实现快速统计，并按照统计得到的次数进行排序，
但是无法保持，当统计内容相同时，其最终结果的排序按照先来后到
的顺序。因此使用map记录索引，vector进行统计的方式可以实现！

```
vector<pair<string, int>>mp;
	map<string, int>con;
	for (auto item : vec)
	{
		string sindex = item.first;
		sindex += " ";
		sindex += to_string(item.second);
        //将两个索引，变成一个进而实现key_value对
		
		if (mp.empty())
		{
			con[sindex] = 0;
			mp.push_back(make_pair(sindex, 1));
		}
		else {
			if (con.count(sindex))
			{
				mp[con[sindex]].second++;
			}
			else {
				con[sindex] = mp.size();
				mp.push_back(make_pair(sindex, 1));
			}
		}
	
	}

```

