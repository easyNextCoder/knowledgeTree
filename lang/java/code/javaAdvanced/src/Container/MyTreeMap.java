package Container;

import java.util.*;
import java.util.Map.Entry;


class Users {  // User对象不再实现任何接口
    private String username;
    private int age;
 
    public Users(String username, int age) {
        super();
        this.username = username;
        this.age = age;
    }
    public String getName() {
    	return this.username;
    }
    
    public int getAge() {
    	return this.age;
    }
}

public class MyTreeMap {
	public static void main(String [] s) {
		Map<Users, String> con = new TreeMap<>(new Comparator<Users>() {
			@Override
			public int compare(Users a, Users b) {
				if(a.getAge() == b.getAge())
					return a.getName().compareTo(b.getName());
				else
					return a.getAge() - b.getAge();
			}
		});
		
		con.put(new Users("jimmy1", 30), "hello");
        con.put(new Users("jimmy2", 30), "hello");
        con.put(new Users("jimmy", 22), "hello");
        con.put(new Users("jimmy", 20), "hello");
	
        for(Entry<Users, String>item:con.entrySet()) {
        	System.out.println(item.getKey() + item.getValue());
        }
	}
}
