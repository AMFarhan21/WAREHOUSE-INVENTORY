'use client'
import { useEffect, useState } from 'react'
import { Stok } from './useGetStokByBarangID'

export interface BarangWithStok {
    id: number
    kode_barang: string
    nama_barang: string
    deskripsi: string
    satuan: string
    harga_beli: number
    harga_jual: number
    created_at: string
    updated_at: string
    stok: Stok | null
}

export interface Meta {
    page: number
    limit: number
    total: number
}


const useGetAllBarangWithStok = () => {
    const [barangsWithStok, setBarangsWithStok] = useState<BarangWithStok[]>([])
    const [meta, setMeta] = useState<Meta | null>(null)
    const [loading, setLoading] = useState(false)
    const [limitNum, setLimitNum] = useState(1000)
    const [pageNum, setPageNum] = useState(0)
    const [search, setSearch] = useState("")
    const [error, setError] = useState("")
    const API_URL = process.env.NEXT_PUBLIC_API_URL


    useEffect(() => {
        const token = localStorage.getItem("TOKEN")
        const getAllBarangWithStok = async () => {
            try {
                setLoading(true)
                const res = await fetch(`${API_URL}/api/barang/stok?search=${search}&page=${pageNum}&limit=${limitNum}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                })

                const data = await res.json()
                if (!res.ok) {
                    setError(data.message)
                    return
                }

                setBarangsWithStok(data.data)
                setMeta(data.meta)
                console.log("Raw API Response:", data) // Cek di console browser, bukan terminal VS Code
                console.log("Cek Stok Item Pertama:", data.data[0]?.stok)
            } catch (error) {

                console.log(error)

            } finally {
                setLoading(false)
            }
        }

        getAllBarangWithStok()
    }, [API_URL, pageNum, search])


    return { barangsWithStok, setBarangsWithStok, loading, error, setPageNum, pageNum, limitNum, setLimitNum, search, setSearch, meta }
}

export default useGetAllBarangWithStok