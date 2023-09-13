import 'package:algolia_helper_flutter/algolia_helper_flutter.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:infinite_scroll_pagination/infinite_scroll_pagination.dart';

class Product {
  final String name;

  Product(this.name);

  static Product fromJson(Map<String, dynamic> json) {
    return Product(json['name']);
  }
}

class HitsPage {
  const HitsPage(this.items, this.pageKey, this.nextPageKey);

  final List<Product> items;
  final int pageKey;
  final int? nextPageKey;

  factory HitsPage.fromResponse(SearchResponse response) {
    final items = response.hits.map(Product.fromJson).toList();
    final isLastPage = response.page >= response.nbPages;
    final nextPageKey = isLastPage ? null : response.page + 1;
    return HitsPage(items, response.page, nextPageKey);
  }
}

final _productsSearcher = HitsSearcher(
  applicationID: dotenv.env["APPLICATION_ID"]!,
  apiKey: dotenv.env["API_KEY"]!,
  indexName: dotenv.env["INDEX_NAME"]!,
);

class SearchMetadata {
  final int nbHits;

  const SearchMetadata(this.nbHits);

  factory SearchMetadata.fromResponse(SearchResponse response) =>
      SearchMetadata(response.nbHits);
}

Stream<SearchMetadata> get _searchMetadata =>
    _productsSearcher.responses.map(SearchMetadata.fromResponse);

Stream<HitsPage> get _searchPage =>
    _productsSearcher.responses.map(HitsPage.fromResponse);

class _SearchViewState extends State<SearchView> {
  final _searchTextController = TextEditingController();
  final PagingController<int, Product> _pagingController =
      PagingController(firstPageKey: 0);

  @override
  void initState() {
    super.initState();
    _searchTextController.addListener(
      () => _productsSearcher.applyState(
        (state) => state.copyWith(
          query: _searchTextController.text,
          page: 0,
        ),
      ),
    );
    _searchPage.listen((page) {
      if (page.pageKey == 0) {
        _pagingController.refresh();
      }
      _pagingController.appendPage(page.items, page.nextPageKey);
    }).onError((error) => _pagingController.error = error);
    _pagingController.addPageRequestListener(
        (pageKey) => _productsSearcher.applyState((state) => state.copyWith(
              page: pageKey,
            )));
  }

  @override
  void dispose() {
    _searchTextController.dispose();
    _productsSearcher.dispose();
    _pagingController.dispose();
    super.dispose();
  }

  Widget _hits(BuildContext context) => PagedListView<int, Product>(
      pagingController: _pagingController,
      builderDelegate: PagedChildBuilderDelegate<Product>(
          noItemsFoundIndicatorBuilder: (_) => const Center(
                child: Text('No results found'),
              ),
          itemBuilder: (_, item, __) => Container(
                color: Colors.white,
                height: 80,
                padding: const EdgeInsets.all(8),
                child: Row(
                  children: [Expanded(child: Text(item.name))],
                ),
              )));

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(title: const Text("Nahrungsmittel")),
        body: Center(
            child: Column(children: [
          SizedBox(
              height: 44,
              child: TextField(
                controller: _searchTextController,
                decoration: const InputDecoration(
                  border: InputBorder.none,
                  hintText: 'Enter a search term',
                  prefixIcon: Icon(Icons.search),
                ),
              )),
          StreamBuilder<SearchMetadata>(
            stream: _searchMetadata,
            builder: (context, snapshot) {
              if (!snapshot.hasData) {
                return const SizedBox.shrink();
              }
              return Padding(
                padding: const EdgeInsets.all(8.0),
                child: Text('${snapshot.data!.nbHits} hits'),
              );
            },
          ),
          Expanded(child: _hits(context))
        ])));
  }
}

class SearchView extends StatefulWidget {
  final int recipeId;
  const SearchView({super.key, required this.recipeId});

  @override
  State<SearchView> createState() => _SearchViewState();

  static const routeName = "/search";
}
