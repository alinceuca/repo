import java.util.Scanner;
public class Jacobi {
	public static double max(double a, double b)
	{
		if(a>b)
			return a;
		return b;
	}

	public static double maximsir(double x[], double y[], int n ,int i)
	{
		if(i==n)
			return x[i]-y[i];
		else
			return max((x[i]-y[i]), maximsir(x,y,n,i+1));
	}
	
	public static double suma(double a[],double x[],int n,int j,int i)
	{
		double s=0;
		for(j=1;j<=n;j++){
			if(j!=i)
				s=s+a[j]*x[j];
		}
		return s;
	}
	
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc=new Scanner(System.in);
		System.out.print("n=");
		int n = sc.nextInt();
		int nrmax =n;
		double norm ;
		int j=1,i=0,nr_it=0,ok=0,eps= 1/1000;
		double a[]= new double[100];
		double y[]= new double[100];
		double b[]= new double[100];
		double x[]= new double[100];
		for(int k=1;k<=n;k++)
			a[k]=sc.nextDouble();
		for(int k=1;k<=n;k++)
			x[k]=sc.nextDouble();
		for(int k=1;k<=n;k++)
			b[k]=sc.nextDouble();
		
		do{
			if(i<n)
				i++;
			y[i]=(b[i]-suma(a,x,n,j,i))/a[i];
			norm=maximsir(x,y,n,i);
			nr_it++;
			x[i]=y[i];
			
		}while(norm>=eps&&nr_it<nrmax);
		
		for(int k=1;k<=n;k++)
			System.out.print(y[k]+" ");
		System.out.println();
		if(norm<eps)
			ok=1;
		if(ok==1)
			System.out.print("Norma e mai mica decat epsilon");
		else
			System.out.print("S-a depasit nr maxim");
		

	}

}


