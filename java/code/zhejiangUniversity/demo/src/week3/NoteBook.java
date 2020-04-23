package week3;
import java.util.ArrayList;
public class NoteBook {

	private ArrayList<String>notes = new ArrayList<String>();
	
	public void add(String s) {
		notes.add(s);
	}
	
	public int getSize() {
		
		return notes.size();
	}
	
	public String getNote(int index) {
		return notes.get(index);
	}
	
	public String removeNote(int index) {
		
		return notes.remove(index);
	}
	
	public static void main(String []s) {
		NoteBook nb = new NoteBook();
		nb.add("first");
		nb.add("second");
		System.out.println(nb.getSize());
	}
}
