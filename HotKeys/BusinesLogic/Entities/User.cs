using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace HotKeysLibrary.Entities
{
    public class User
    {
        public Guid Id;
        public string Nickname;
        private List<Program> programs;

        public User(Guid id, string nickname, List<Program> programs)
        {
            Id = id;
            Nickname = nickname;
            this.programs = programs;
        }

        public IEnumerable<Program> Programs => programs.AsEnumerable();

    }
}
