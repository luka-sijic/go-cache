import ProductList from './ProductList';

async function getProducts() {
  const res = await fetch('http://localhost:8086/products', {
    cache: 'no-store',
  });
  return res.json();
}

export default async function ProductsPage() {
    const products = await getProducts();

    return <ProductList products={products} />; 
}
