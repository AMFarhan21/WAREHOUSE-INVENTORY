'use client'

import { useAuth } from "@/context/AuthContext";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function RootPage() {
  const router = useRouter()
  const { token, loading } = useAuth()

  useEffect(() => {
    if (!loading) {
      if (token) {
        router.replace("/dashboard")
      } else {
        router.replace("/auth/login")
      }
    }
  }, [token, loading, router])

  return (
    <div>Loading...</div>
  );
}
