package week4;

import java.util.Scanner;
public class Item {
	//private String title;
	public static void main(String [] args) {
		Scanner input = new Scanner(System.in);
		double result = 0;
		while(true) {
			double a = input.nextDouble();
			double b = input.nextDouble();
			result = a*b+result;
			System.out.println(result);
		}
	}
}
