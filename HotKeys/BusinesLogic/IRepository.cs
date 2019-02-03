using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BisinessLogic
{
    interface IRepository<T> where T: class
    {
        T Get(Guid id);
        void Set(Guid id, T item);
    }
}
