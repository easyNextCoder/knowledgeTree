package LRUCache;

import java.util.LinkedHashMap;
import java.util.Map;

class LRUCaches<K, V> extends LinkedHashMap<K, V>{
	private int maxEntries;
	
	public LRUCaches(int maxEntries) {
		super(16, 0.75f, true);
		this.maxEntries = maxEntries;
	}
	
	@Override
	protected boolean removeEldestEntry(Map.Entry<K, V> eldest) {
		return size() > maxEntries;
	}
}

public class LRUCacheTest {
	public static void main(String []s) {
		LRUCaches<String, Object>cache = new LRUCaches<>(3);
		cache.put("a","abstract");
        cache.put("b","basic");
        cache.put("c","call");
        cache.get("a");
        cache.put("d","µÎµÎµÎ");
        System.out.println(cache);
	}
}
