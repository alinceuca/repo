import java.util.Scanner;
public class SistemTridiagonalRecursiv {

	static Scanner sc = new Scanner(System.in);
	static double a[],b[],c[],d[],r[],s[],x[];
	static int n=sc.nextInt();
	
	public static void citire() {
		for(int i=1;i<=n;i++)
			a[i]=sc.nextDouble();

		for(int i=1;i<=n-1;i++)
			b[i]=sc.nextDouble();

		for(int i=1;i<=n-1;i++)
			c[i]=sc.nextDouble();

		for(int i=1;i<=n;i++)
			d[i]=sc.nextDouble();
	}
	
	public static void rs(int i) {
		if(i==1 && a[1]!=0) {
			r[1]=-c[1]/a[1];
			s[1]=d[1]/a[1];
		}
		if(i<n-1) {
			r[i+1]= -c[i+1]/(b[i] * r[i] + a[i+1]);
			s[i+1]= (d[i+1]-b[i] * s[i]) / (b[i] * r[i] + a[i+1]);
			rs(i+1);
		}
	}

	public static void sol(int i) {
		if (i==n) {
			x[i]=(d[i]-b[i-1]*s[i-1])/(b[i-1]*r[i-1]+a[i]);
			sol(i-1);
		} else if (i > 0) {
			x[i] = r[i] * x[i + 1] + s[i];
			sol(i-1);
		}
	}

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		a = new double[n+1];
		b = new double[n];
		c = new double[n];
		d = new double[n+1];
		r = new double[n];
		s = new double[n];
		x = new double[n+1];	
		
		citire();

		rs(1);
		sol(n);

		System.out.println("Solutiile sunt: ");
		for (int i = 1; i <= n; i++)
			System.out.println("x[" + i + "]: " + (double) Math.round(x[i] * 100000d) / 100000d);
		
	}
}
