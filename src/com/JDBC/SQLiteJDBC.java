package com.JDBC;
import java.sql.*;
public class SQLiteJDBC {
    private static String myurl = "jdbc:sqlite:insuranceDB.db";
    //private static String user = "root";
    //private static String password = "";
    //private static String user = "remote_user";
    //private static String password = "remote";
    private static String driverClass = "org.sqlite.JDBC";
    private static Connection conn;

    public static Connection getConnection(){

        try {
            Class.forName(driverClass);
        } catch (ClassNotFoundException error) {
            System.err.print("Driver Not Found!");
        }

        try {
            //conn = DriverManager.getConnection(myurl);
            conn = DriverManager.getConnection("jdbc:sqlite:insuranceDB.db");
            conn.setAutoCommit(false);

            return conn;
        } catch (SQLException error) {
            error.printStackTrace();
            throw new RuntimeException(error);
        }
    }

    public static void close(Connection conn, Statement stat) {
        if (stat != null) {
            try {
                stat.close();
            } catch (SQLException e) {
                int i;
                e.printStackTrace();
                throw new RuntimeException(e);
            }
        }
        if (conn != null) {
            try {
                conn.close();
            } catch (SQLException e) {
                e.printStackTrace();
                throw new RuntimeException(e);
            }
        }
    }

    public static void close(Connection conn, Statement stat, ResultSet results) {
        if (results != null)
            try {
                results.close();
            } catch (SQLException e) {
                e.printStackTrace();
                throw new RuntimeException(e);
            }
        close(conn, stat);
    }

    public static void main( String args[] )
    {
        Connection c = null;
        Statement stmt = null;
        try {
            c=getConnection();
            System.out.println("Opened database successfully");
            stmt = c.createStatement();
            // String sql = "DELETE from produce where rowid=2;";
            //  stmt.executeUpdate(sql);
            //  c.commit();

            ResultSet rs = stmt.executeQuery( "SELECT * FROM produce;" );
            while ( rs.next() ) {
                //int id = rs.getInt("id");
                String  name = rs.getString("name");
                int age  = rs.getInt("rank");
                int  address = rs.getInt("money");
                float salary = rs.getFloat("score");
                //System.out.println( "rowid = " + id );
                System.out.println( "name = " + name );
                System.out.println( "rank = " + age );
                System.out.println( "money = " + address );
                System.out.println( "score = " + salary );
                System.out.println();
            }
            rs.close();
            stmt.close();
            c.close();
        } catch ( Exception e ) {
            System.err.println( e.getClass().getName() + ": " + e.getMessage() );
            System.exit(0);
        }
        System.out.println("Operation done successfully");
        close(c,stmt);
    }
}
