import type { Metadata } from "next";
import "./globals.css";
import Header from "@/app/ui/Header";
import Footer from "@/app/ui/Footer";

export const metadata: Metadata = {
  title: "ToDo app with Go backend",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body className="bg-zinc-700 text-white">
        <Header />
        <div className="p-2 min-h-screen">{children}</div>
        <Footer />
      </body>
    </html>
  );
}
