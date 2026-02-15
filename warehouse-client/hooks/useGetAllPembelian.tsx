'use client'
import { useEffect, useState } from 'react'

export interface MasterBarang {
    id: number
    kode_barang: string
    nama_barang: string
    deskripsi: string
    satuan: string
    harga_beli: number
    harga_jual: number
    created_at: string
    updated_at: string
}

export interface Meta {
    page: number
    limit: number
    total: number
}


const useGetAllPembelian = () => {
    const [pembelian, setPembelian] = useState<MasterBarang[]>([])
    const [meta, setMeta] = useState<Meta | null>(null)
    const [loading, setLoading] = useState(false)
    const [limitNum, setLimitNum] = useState(8)
    const [pageNum, setPageNum] = useState(0)
    const [search, setSearch] = useState("")
    const [error, setError] = useState("")
    const API_URL = process.env.NEXT_PUBLIC_API_URL


    useEffect(() => {
        const token = localStorage.getItem("TOKEN")
        const getAllBarang = async () => {
            try {
                setLoading(true)
                const res = await fetch(`${API_URL}/api/barang?search=${search}&page=${pageNum}&limit=${limitNum}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                })

                const data = await res.json()
                if (!res.ok) {
                    setError(data.message)
                    return
                }

                setPembelian(data.data)
                setMeta(data.meta)
            } catch (error) {

                console.log(error)

            } finally {
                setLoading(false)
            }
        }

        getAllBarang()
    }, [API_URL, pageNum, search])


    return { pembelian, setPembelian, loading, error, setPageNum, pageNum, limitNum, search, setSearch, meta }
}

export default useGetAllPembelian