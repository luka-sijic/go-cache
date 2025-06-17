'use client';

import { useEffect, useState } from "react";

type Product = {
  id: number;
  name: string;
  price: number;
};

export default function ProductList({ products }: { products: Product[] }) {
    const [editableProducts, setEditableProducts] = useState(() => [...products]);


    function handleClick(product: Product) {

        console.log('Clicked', product);
    }

    function handlePriceChange(id: number, newPrice: number) {
        const updated = editableProducts.find(p => p.id === id);
        if (!updated) return;

        setEditableProducts((prev) => 
           prev.map((p) => (p.id === id ? {...p, price: newPrice } : p))
        );

        fetch('http://localhost:8086/products', {
           method: 'PATCH',
           headers: {
               'Content-Type': 'application/json',
           },
           body: JSON.stringify({ id, name: updated.name, price: newPrice }),
        });
    }
    
    useEffect(() => {
        const evtSource = new EventSource('http://localhost:8086/events');
        evtSource.addEventListener('update', (e) => {
            const data: Product = JSON.parse(e.data);
            setEditableProducts(prev =>
                prev.map(p => (p.id === data.id ? data : p))
            );
        });

        return () => {
            evtSource.close();
        };
    }, []);

    return (
    <ul className="space-y-4">
      {editableProducts.map((product) => (
        <li key={product.id} className="flex justify-between border p-4">
          <div>
            <p>{product.name}</p>
            <input
                type='number'
                value={product.price}
                onChange={(e) => {
                    const newPrice = parseFloat(e.target.value);
                    setEditableProducts(prev =>
                        prev.map(x =>
                            x.id === product.id ? { ...x, price: newPrice } : x
                        )
                    );
                }}
                onKeyDown={e => {
                    if (e.key === 'Enter') {
                        const price = parseFloat((e.target as HTMLInputElement).value)
                        handlePriceChange(product.id, price);
                    }
                }}
            />
          </div>
          <button
            onClick={() => handleClick(product)}
            className="bg-blue-600 text-white px-4 py-2 rounded"
          >
            Select
          </button>
        </li>
      ))}
    </ul>
  );
}

