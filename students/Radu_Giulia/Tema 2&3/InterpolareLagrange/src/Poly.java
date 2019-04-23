import java.util.Scanner;

public class Poly {

	public static int x[],y[];
	public static int n;
	public static double z;
	
	static Scanner sc=new Scanner(System.in);
	
	static double res()
	{
		int s=0,p=1;
		for (int i = 1; i <=n; i++) {
	        for (int j = 1; j <= n; j++) {
	            if (j != i) 
	                p *= (z - x[j]) / (x[i] - x[j]);
	        }
	        s += p * y[i];
	        p = 1;  
	    }
		return s;
	}
	
	
	public static void main(String[] args) {
		
		System.out.println("N:");
		n=sc.nextInt();
		
		x=new int[n+2];
		y=new int[n+2];
		
		System.out.println("Perechi de coordonate");
		
		for(int i=1;i<=n;i++) {
			x[i]=sc.nextInt();
			y[i]=sc.nextInt();
		}
		
		System.out.println("z pentru L(z) ?:  ");
		z=sc.nextDouble();
		
		System.out.println("L(z)= " + res());

	}

}
