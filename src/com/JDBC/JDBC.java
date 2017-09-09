package com.JDBC;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
/**
 * Created by hasee on 2017/5/4.
 */
public class JDBC {
    private static String myurl = "jdbc:mysql://127.0.0.1/hhlab?" +
            "useUnicode=true" +
            "&characterEncoding=UTF-8" +
            "&zeroDateTimeBehavior=convertToNull";
    //private static String user = "root";
    //private static String password = "";
    private static String user = "remote_user";
    private static String password = "remote";
    private static String driverClass = "com.mysql.jdbc.Driver";
    private static Connection conn;

    public static Connection getConnection(){

        try {
            Class.forName(driverClass);
        } catch (ClassNotFoundException error) {
            System.err.print("Driver Not Found!");
        }

        try {
            conn = DriverManager.getConnection(myurl, user, password);
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
}
