using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BisinessLogic
{
    public class HotKey
    {
        public Guid Id { get; }
        public string Description { get; }
        private List<Key> Combination;
        public IEnumerable<Key> combination => Combination.AsEnumerable();


        public HotKey(Guid id, string description, List<Key> combination)
        {
            Id = id;
            Description = description;
            Combination = combination;
        }
    }
}
