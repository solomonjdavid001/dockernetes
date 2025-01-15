"use client";

import { DataTable } from "../../components/datatable/data-table";
import axios from "axios";
import { columns, Image } from "./columns";
import { useEffect, useState } from "react";

export default function Payments() {
  const [data, setData] = useState<Image[]>([]);

  useEffect(() => {
    axios
      .get("http://localhost:8080/api/v1/images")
      .then((response) => {
        console.log("API Response:", response.data); // Log the full response
        setData(response.data.images || []); // Make sure the imageList exists
      })
      .catch((error) => {
        console.error("Error fetching data:", error);
      });
  }, []);

  return (
    <div className="container mx-auto py-6">
      <h1 className="font-bold text-2xl pb-10">Images</h1>
      <DataTable columns={columns} data={data} />
    </div>
  );
}
