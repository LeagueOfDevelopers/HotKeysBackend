using System;

namespace HotKeysLibrary.Infrastructure
{
    public interface IRepository<T> where T : class
    {
        T Get(Guid id);
        void Add(Guid id, T item);
    }
}
