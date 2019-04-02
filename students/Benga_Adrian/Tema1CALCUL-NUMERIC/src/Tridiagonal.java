import java.util.Scanner;
public class Tridiagonal {
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc=new Scanner(System.in);
		System.out.print("n=");
		int n = sc.nextInt();
		int a[]= new int[100];
		int b[]= new int[100];
		int c[]= new int[100];
		int d[]= new int[100];
		double x[]= new double[100];
		int r[]= new int[100];
		int s[]= new int[100];
		for(int k=1;k<=n;k++)
			a[k]=sc.nextInt();
		for(int k=2;k<n;k++)
			b[k]=sc.nextInt();
		for(int k=1;k<=n-1;k++)
			c[k]=sc.nextInt();
		for(int k=1;k<=n;k++)
			d[k]=sc.nextInt();
		r[1]= -c[1]/a[1];
		r[1]=d[1]/a[1];
		for(int k=2;k<n-1;k++){
			r[k]= -c[k]/(b[k]*r[k-1]+a[k-1]);
			s[k]= (d[k]-b[k]*s[k-1])/(b[k]*r[k-1]+a[k]);
			
		}
		
		for(int k=n;k>=2;k--)
			x[k]=(d[k]-b[k]*s[k-1])/(b[k]*r[k-1]+a[k]);
		x[1]=(-c[1]*c[2])/a[1]+d[1]/a[1];
	
		for(int k=0;k<=n;k++)
			System.out.print(x[k]+" ");
	}	

}

