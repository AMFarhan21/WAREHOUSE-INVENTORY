import { useRouter } from 'next/navigation'
import React, { useEffect, useState } from 'react'

export interface HistoryStok {
    id: number
    barang_id: number
    user_id: number
    jenis_transaksi: string
    jumlah: number
    stok_sebelum: number
    stok_sesudah: number
    keterangan: string
    created_at: string
    barang: Barang
    user: User
}

export interface Barang {
    id: number
    kode_barang: string
    nama_barang: string
    satuan: string
    harga_jual: number
}

export interface User {
    id: number
    username: string
    full_name: string
}

interface Meta {
    page: number
    limit: number
    total: number
}

const useGetHistoryByBarangID = (barangID: number) => {
    const [historyStoksByBarangID, setHistoryStokByBarangID] = useState<HistoryStok[]>([])


    const [meta, setMeta] = useState<Meta | null>(null)
    const [loading, setLoading] = useState(false)
    const [limitNum, setLimitNum] = useState(5)
    const [pageNum, setPageNum] = useState(0)
    const [error, setError] = useState("")
    const router = useRouter()
    const API_URL = process.env.NEXT_PUBLIC_API_URL

    useEffect(() => {
        const token = localStorage.getItem("TOKEN")
        const getAllBarang = async () => {
            try {
                setLoading(true)
                const res = await fetch(`${API_URL}/api/history-stok/${barangID}?page=${pageNum}&limit=${limitNum}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                })

                const data = await res.json()
                if (data.error_code == "UNAUTHORIZED") {
                    localStorage.removeItem("TOKEN")
                    router.replace("/auth/login")
                    return
                }



                if (!res.ok) {
                    setError(data.message)
                    return
                }

                setHistoryStokByBarangID(data.data)
                setMeta(data.meta)
            } catch (error) {

                console.log(error)

            } finally {
                setLoading(false)
            }
        }

        getAllBarang()
    }, [API_URL, pageNum])


    return { historyStoksByBarangID, loading, error, setPageNum, pageNum, limitNum, meta }

}

export default useGetHistoryByBarangID