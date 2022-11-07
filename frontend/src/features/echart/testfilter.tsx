  import React, { useCallback,useEffect, useState } from 'react'
  
  function ProductFilters(props) {
    const { 
      categories,
      onFilterChange,
    } = props
    
    return (
      <section 
        className="filters"
        aria-labelledby="filters-header">
        <header id="filters-header">
          {'Filters'}
        </header>
        
        <ul>
          {categories.map(category => (
            <li key={category}>
              <label>
                <input 
                  onChange={onFilterChange}
                  type="checkbox"
                  value={category} />
                {category}
              </label>
            </li>
          ))}
        </ul>
      </section>
    )
  }
  
  function Product(props) {
    const { product } = props
    
    return (
      <li
        key={product.SummaryId}
        className="product">
        <div className="product-details">
          <header>{product.stationName}</header>
          <div className="category">{product.year}</div>
        </div>
      </li>
    )
  }
  
  function ProductsList(props) {
    const { products } = props
    
    return (
      <ul className="products">
        {products.map(product => (
          <Product product={product} />
        ))}
      </ul>
    )
  }
  
  export const Testcheckbox = ({filter,datax}:any) => {
    console.log(datax)
    const [state, setState] = useState({
      products: datax,
      filters: new Set(),
    })
    
    const handleFilterChange = useCallback(event => {
      setState(previousState => {
        let filters = new Set(previousState.filters)
        let products = datax
        
        if (event.target.checked) {
          filters.add(event.target.value)
        } else {
          filters.delete(event.target.value)
        }
        
        if (filters.size) {
          products = products.filter(product => {
            return filters.has(product.category)
          })
        }
        
        return {
          filters,
          products,
        }
      })
    }, [setState])
    
    return (
      <main>
        <ProductFilters 
          categories={filter}
          onFilterChange={handleFilterChange}/>
        <ProductsList products={state.products} />
      </main>
    )
  }
