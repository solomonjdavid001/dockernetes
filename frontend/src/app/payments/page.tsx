"use client";

import { DataTable } from "./data-table";
import axios from "axios";
import { columns, Image } from "./columns";
import { useEffect, useState } from "react";

export default function Payments() {
  const [data, setData] = useState<Image[]>([]);
  const [loading, setLoading] = useState(true); // Loading state

  useEffect(() => {
    axios
      .get("http://localhost:8080/api/v1/images")
      .then((response) => {
        console.log("API Response:", response.data); // Log the full response
        setData(response.data.images || []); // Make sure the imageList exists
      })
      .catch((error) => {
        console.error("Error fetching data:", error);
      })
      .finally(() => {
        setLoading(false); // Stop loading after the request completes
      });
  }, []);

  if (loading) {
    return <div>Loading...</div>; // Show loading state
  }

  return (
    <div className="container mx-auto py-10">
      <DataTable columns={columns} data={data} />
    </div>
  );
}
