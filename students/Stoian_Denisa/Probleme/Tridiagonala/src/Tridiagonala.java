import java.util.Scanner;

public class Tridiagonala {


	public static void rezolvare(double[] b, double[] a, double[] c, double[] d, int n) {

	double[]R=new double[n]; double[]S=new double[n]; double[]x=new double[n];
		R[0]=(-c[1])/a[0];    
		S[0]=d[0]/a[0];
		 
	for(int i=1;i<n-1;i++)
		R[i]=(-c[i])/(a[i]+b[i]*R[i-1]);
		 
	for(int i=1;i<n-1;i++)
		S[i]=(d[i]-b[i]*S[i-1])/(a[i]+b[i]*R[i-1]);
		 
	for(int i=1;i<n;i++)x[i]=(d[i]-b[i]*S[i-1])/(b[i]*R[i-1]+a[i]);
	            x[0]=R[0]*x[1]+S[0];
		
		for(int i=0;i<n;i++)System.out.print(x[i]+" ");
			
		}
	
	static Scanner sc = new Scanner (System.in);
	public static void main(String[] args) {
	System.out.println("n=");
	int n=sc.nextInt();
	System.out.print("a=");
	double[]a=new double[n];
		for(int i=0;i<n;i++)
			a[i]=sc.nextInt();
	System.out.print("c=");
    double[]c=new double[n];
		for(int i=1;i<n;i++)
			c[i]=sc.nextInt();
	double[]b=new double[n];
	System.out.print("b=");
		for(int i=0;i<n-1;i++)
		b[i]=sc.nextInt();
	System.out.print("d=");
    double[]d=new double[n];
		for(int i=0;i<n;i++)
		d[i]=sc.nextInt();

  rezolvare(b,a,c,d,n);

	 }
}
