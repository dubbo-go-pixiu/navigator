package trie

import (
	"container/list"
	"navigator/pkg/algorithm/heap"
)

type IdentityPath[T any] struct {
	list.List
}

func (path IdentityPath[T]) Peek() T {
	if path.Len() == 0 {
		return nil
	}
	return path.Front().Value
}

func New[T any]() IdentityPath[T] {
	var ret = IdentityPath[T]{}
	ret.Init()
	return ret
}

func FromSorted[T any](sorted heap.Sorted[T]) IdentityPath[T] {
	var ret = IdentityPath[T]{}
	ret.Init()
	for sorted.Len() > 0 {
		ret.PushBack(sorted.Pop())
	}
	return ret
}

//public  static <I> IdentityPath<I> of(List<I> ids){
//IdentityPath<I> path = new IdentityPath();
//ids.stream().forEach(i -> path.append(i));
//return path;
//}
//
//public  static <I> IdentityPath<I> of(I[] ids){
//IdentityPath<I> path = new IdentityPath();
//Arrays.stream(ids).forEach(i -> path.append(i));
//return path;
//}
//LinkedList<T> pathList = new LinkedList<>();
//
//public static IdentityPath<Character> of(char[] chars) {
//IdentityPath<Character> path = new IdentityPath();
//for (int i=0;i<chars.length;i++){
//path.append(chars[i]);
//}
//return path;
//}
//
//public IdentityPath<T> append(T path){
//pathList.add(path);
//return this;
//}
//public IdentityPath<T> push(T path){
//pathList.add(0,path);
//return this;
//}
//public IdentityPath<T> appendAll(IdentityPath<T> path){
//pathList.addAll(path.pathList);
//return this;
//}
//
////    public static IdentityPath of(String path){
////        IdentityPath p = new IdentityPath();
////        String[] pArr = path.split(".");
////        p.pathList.addAll( Arrays.stream(pArr).collect(Collectors.toList()));
////        return p;
////    }
//
//public T top(){
//if (pathList.size()==0){
//return null;
//}
//return pathList.get(0);
//}
//
//public T pop(){
//if (pathList.size()==0){
//return null;
//}
//return pathList.remove(0);
//}
//
//public T peek(){
//if (pathList.size()==0){
//return null;
//}
//return pathList.get(0);
//}
//public IdentityPath<T> deepCopy(){
//IdentityPath<T> ret = new IdentityPath();
//ret.pathList.addAll(this.pathList);
//return ret;
//}
//
//public int size(){
//return pathList.size();
//}
//
//@Override
//public String toString() {
//return pathList.stream().map(k->k.toString()).collect(Collectors.joining("."));
//}
