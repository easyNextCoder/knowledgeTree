package week4;

import test.Item;

public class DVD extends Item {
	private String director;
	private int playingTime;
	private boolean gotIt = false;
	private String comment;
	public DVD(String director, int playingTime, boolean gotIt, String comment) {
		super();
		this.director = director;
		this.playingTime = playingTime;
		this.gotIt = gotIt;
		this.comment = comment;
	}
	
	
}
