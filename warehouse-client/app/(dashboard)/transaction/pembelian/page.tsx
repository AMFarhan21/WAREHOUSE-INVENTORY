"use client"

import { Checkbox } from "@/components/ui/checkbox"
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"
import useGetAllBarangWithStok from "@/hooks/useGetAllBarangWithStok"
import { Input } from "@/components/ui/input"
import { useState } from "react"
import useCreatePembelian from "@/hooks/useCreatePembelian"
import { Label } from "@/components/ui/label"
import { Button } from "@/components/ui/button"
import toast from "react-hot-toast"

const Page = () => {
    const { barangsWithStok } = useGetAllBarangWithStok()
    const { createPembelian, loading, error } = useCreatePembelian()
    const [supplier, setSupplier] = useState("")
    const [qtys, setQtys] = useState<Record<string, number>>({})


    const handleQtyChange = (id: string, value: number) => {
        setQtys(prev => ({
            ...prev,
            [id]: value
        }))
    }

    const [selectedRows, setSelectedRows] = useState<Set<string>>(
        new Set([""])
    )

    const selectAll = selectedRows.size === barangsWithStok.length

    const handleSelectAll = (checked: boolean) => {
        if (checked) {
            setSelectedRows(new Set(barangsWithStok.map((item) => String(item.id))))
        } else {
            setSelectedRows(new Set())
        }
    }

    const handleSelectRow = (id: string, checked: boolean) => {
        const newSelected = new Set(selectedRows)
        if (checked) {
            newSelected.add(id)
        } else {
            newSelected.delete(id)
        }
        setSelectedRows(newSelected)
    }

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        if (!supplier) {
            toast.error("Isi nama supplier")
            return
        }

        if (selectedRows.size == 0) {
            toast.error("Pilih minimal satu barang")
            return
        }

        if (error) {
            toast.error(error)
            return
        }

        const beli_detail = Array.from(selectedRows).filter(id => id !== "").map(id => ({
            barang_id: Number(id),
            qty: qtys[id] ?? 1
        }))

        console.log("Payload yang dikirim:", { supplier, beli_detail });
        const res = await createPembelian(supplier, beli_detail)

        if (res) {
            toast.success("Pembelian berhasil")
        }

        setSupplier("")
        setSelectedRows(new Set())
        setQtys({})
    }

    return (
        <div className="w-full min-h-screen flex">
            <div className='w-300 mx-auto justify-center'>
                <div className='text-3xl font-bold mt-12'>PEMBELIAN</div>
                <div className="mt-8 flex w-full items-center justify-between">
                    <div className="w-120 flex flex-col gap-2">
                        <Label>Nama Supplier</Label>
                        <Input name="supplier" value={supplier} onChange={e => setSupplier(e.target.value)} placeholder="PT Contoh Industri" />

                    </div>
                    <Button className="mt-5" type="submit" onClick={handleSubmit} disabled={loading}>
                        {loading ? "Processing..." : "Save Purchase"}
                    </Button>
                </div>

                <Table className="mt-4">
                    <TableHeader>
                        <TableRow>
                            <TableHead className="w-8">
                                <Checkbox
                                    id="select-all-checkbox"
                                    name="select-all-checkbox"
                                    checked={selectAll}
                                    onCheckedChange={handleSelectAll}
                                />
                            </TableHead>
                            <TableHead>Nama Barang</TableHead>
                            <TableHead>Kode Barang</TableHead>
                            <TableHead>Harga Beli</TableHead>
                            <TableHead>Harga Jual</TableHead>
                            <TableHead>Stok Akhir</TableHead>
                            <TableHead>Qty</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {barangsWithStok.map((item) => (
                            <TableRow
                                key={item.id}
                                data-state={selectedRows.has((String(item.id))) ? "selected" : undefined}
                            >
                                <TableCell>
                                    <Checkbox
                                        id={`row-${item.id}-checkbox`}
                                        name={`row-${item.id}-checkbox`}
                                        checked={selectedRows.has(String(item.id))}
                                        onCheckedChange={(checked) =>
                                            handleSelectRow(String(item.id), checked === true)
                                        }
                                    />
                                </TableCell>
                                <TableCell className="font-medium">{item.nama_barang}</TableCell>
                                <TableCell className="font-medium">{item.kode_barang}</TableCell>
                                <TableCell>{item.harga_beli}</TableCell>
                                <TableCell>{item.harga_jual}</TableCell>
                                <TableCell>{item.stok ? item.stok.stok_akhir : 0}</TableCell>
                                <TableCell className="w-30"><Input name="qty" value={qtys[item.id] || 1} onChange={e => handleQtyChange(String(item.id), e.target.valueAsNumber)} type="number" min={1} required /></TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </div>
        </div>
    )
}

export default Page