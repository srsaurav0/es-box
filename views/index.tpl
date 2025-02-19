<!DOCTYPE html>
<html>
<head>
    <title>Beego</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <script>
        async function fetchSuggestions(query) {
            if (query.length === 0) {
                document.getElementById("suggestions").innerHTML = '';
                return;
            }
            try {
                const response = await fetch('/search', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ query: query }),
                });
                const data = await response.json();
                if (response.ok) {
                    displaySuggestions(data);
                } else {
                    console.error("Error fetching suggestions:", data.error);
                }
            } catch (error) {
                console.error("Error:", error);
            }
        }

        function displaySuggestions(suggestions) {
            const suggestionList = document.getElementById("suggestions");
            suggestionList.innerHTML = '';
            
            if (suggestions.length > 0) {
                const gridContainer = document.createElement("div");
                gridContainer.classList.add("grid", "grid-cols-2", "gap-2", "p-2");
                
                suggestions.forEach((suggestion) => {
                    const suggestionItem = document.createElement("div");
                    suggestionItem.classList.add(
                        "px-3",
                        "py-2",
                        "text-gray-700",
                        "hover:bg-gray-200",
                        "cursor-pointer",
                        "rounded",
                        "truncate"
                    );
                    suggestionItem.textContent = suggestion.name;
                    suggestionItem.onclick = function() {
                        document.getElementById("default-search").value = suggestion.name;
                        document.getElementById("suggestions").innerHTML = '';
                        displayProductCard(suggestion);
                    };
                    gridContainer.appendChild(suggestionItem);
                });
                suggestionList.appendChild(gridContainer);
            } else {
                suggestionList.innerHTML = '<div class="px-4 py-2 text-gray-500">No suggestions found</div>';
            }
        }

        function displayProductCard(product) {
            const productCard = document.getElementById("product-card");
            productCard.innerHTML = `
                <div class="bg-white rounded-lg shadow-md p-6">
                    <h2 class="text-xl font-semibold mb-2">${product.name}</h2>
                    <p class="text-gray-600">Price: $${product.price.toFixed(2)}</p>
                </div>
            `;
            productCard.classList.remove('hidden');
        }
    </script>
</head>
<body class="bg-gray-50">
    <div class="container mx-auto px-4">
        <div class="flex justify-center relative">
            <!-- Search Form -->
            <div class="w-2/5 mt-20">
                <form>
                    <label for="default-search" class="mb-2 text-sm font-medium text-gray-900 sr-only">Search</label>
                    <div class="relative">
                        <div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                            <svg class="w-4 h-4 text-gray-500" aria-hidden="true" xmlns="http://www.w3.org/2000/svg"
                                fill="none" viewBox="0 0 20 20">
                                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z" />
                            </svg>
                        </div>
                        <input type="search" id="default-search"
                            class="block w-full p-4 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-white focus:ring-blue-500 focus:border-blue-500"
                            placeholder="Enter Product Name" required oninput="fetchSuggestions(this.value)" />
                        
                        <div id="suggestions" class="absolute w-full bg-white shadow-lg border border-gray-300 mt-1 rounded-lg z-10"></div>

                        <button type="submit"
                            class="text-white absolute end-2.5 bottom-2.5 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2">
                            Search
                        </button>
                    </div>
                </form>
            </div>

            <!-- Product Card -->
            <div id="product-card" class="hidden absolute right-0 mt-20 w-1/4"></div>
        </div>
    </div>
</body>
</html>